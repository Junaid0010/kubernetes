/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package client

import (
	"context"
	"errors"
	"io"
	"math/rand"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
	"k8s.io/klog/v2"
	"sigs.k8s.io/apiserver-network-proxy/konnectivity-client/proto/client"
)

// Tunnel provides ability to dial a connection through a tunnel.
type Tunnel interface {
	// Dial connects to the address on the named network, similar to
	// what net.Dial does. The only supported protocol is tcp.
	Dial(protocol, address string) (net.Conn, error)

	// Close closes a GRPC Tunnel
	Close()
}

type dialResult struct {
	err    string
	connid int64
}

// grpcTunnel implements Tunnel
type grpcTunnel struct {
	stream          client.ProxyService_ProxyClient
	pendingDial     map[int64]chan<- dialResult
	conns           map[int64]*conn
	pendingDialLock sync.RWMutex
	connsLock       sync.RWMutex
	streamLock      sync.RWMutex
	singleUse       bool
	proxyConn       *grpc.ClientConn
	stopCh          chan struct{}
	dialed          chan struct{}
}

// CreateReusableGrpcTunnel creates a Tunnel to dial to a remote server through a
// gRPC based proxy service.
// The Dial() method of the returned tunnel can be called multiple times.
// The tunnel must be closed by calling Tunnel.Close()
func CreateReusableGrpcTunnel(address string, opts ...grpc.DialOption) (Tunnel, error) {
	return createGRPCTunnel(false, address, opts...)
}

// CreateSingleUseGrpcTunnel creates a Tunnel to dial to a remote server through a
// gRPC based proxy service.
// Currently, a single tunnel supports a single connection, and the tunnel is closed when the connection is terminated
// The Dial() method of the returned tunnel should only be called once
// The tunnel will automatically close after the initial connection created by the Dial function is closed
func CreateSingleUseGrpcTunnel(address string, opts ...grpc.DialOption) (Tunnel, error) {
	return createGRPCTunnel(true, address, opts...)
}

func createGRPCTunnel(singleUse bool, address string, opts ...grpc.DialOption) (Tunnel, error) {
	c, err := grpc.Dial(address, opts...)
	if err != nil {
		return nil, err
	}

	grpcClient := client.NewProxyServiceClient(c)

	stream, err := grpcClient.Proxy(context.Background())
	if err != nil {
		return nil, err
	}

	tunnel := &grpcTunnel{
		stream:      stream,
		pendingDial: make(map[int64]chan<- dialResult),
		conns:       make(map[int64]*conn),
		singleUse:   singleUse,
		proxyConn:   c,
		stopCh:      make(chan struct{}),
		dialed:      make(chan struct{}),
	}

	go tunnel.serve()

	return tunnel, nil
}

func (t *grpcTunnel) serve() {
	defer t.proxyConn.Close()

	for {
		pkt, err := t.stream.Recv()
		if err == io.EOF {
			klog.Warning("Warning: EOF, closing")
			return
		}
		if err != nil || pkt == nil {
			klog.Warningf("stream read error: %v", err)
			return
		}

		klog.V(6).Infof("[tracing] recv packet, type: %s", pkt.Type)

		switch pkt.Type {
		case client.PacketType_DIAL_RSP:
			resp := pkt.GetDialResponse()
			t.pendingDialLock.RLock()
			ch, ok := t.pendingDial[resp.Random]
			t.pendingDialLock.RUnlock()

			if !ok {
				klog.Warning("DialResp not recognized; dropped")
			} else {
				ch <- dialResult{
					err:    resp.Error,
					connid: resp.ConnectID,
				}
			}
		case client.PacketType_DATA:
			resp := pkt.GetData()
			// TODO: flow control
			t.connsLock.RLock()
			conn, ok := t.conns[resp.ConnectID]
			t.connsLock.RUnlock()

			if ok {
				conn.readCh <- resp.Data
			} else {
				klog.Warningf("connection id %d not recognized", resp.ConnectID)
			}
		case client.PacketType_CLOSE_RSP:
			resp := pkt.GetCloseResponse()
			t.connsLock.RLock()
			conn, ok := t.conns[resp.ConnectID]
			t.connsLock.RUnlock()

			if ok {
				close(conn.readCh)
				conn.closeCh <- resp.Error
				close(conn.closeCh)
				t.connsLock.Lock()
				delete(t.conns, resp.ConnectID)
				t.connsLock.Unlock()
				if t.singleUse {
					t.Close()
				}
			}
			klog.Warningf("connection id %d not recognized", resp.ConnectID)
		}
		select {
		case <-t.stopCh:
			return
		default:
		}
	}
}

// Dial connects to the address on the named network, similar to
// what net.Dial does. The only supported protocol is tcp.
func (t *grpcTunnel) Dial(protocol, address string) (net.Conn, error) {
	if t.singleUse {
		select {
		case <-t.dialed:
			klog.Warningf("Dialing multiple times is not permitted on singleUse GRPC Tunnels")
			return nil, errors.New("Dialing multiple times is not permitted on singleUse GRPC Tunnels")
		default:
			close(t.dialed)
		}
	}

	if protocol != "tcp" {
		return nil, errors.New("protocol not supported")
	}

	random := rand.Int63()
	resCh := make(chan dialResult)
	t.pendingDialLock.Lock()
	t.pendingDial[random] = resCh
	t.pendingDialLock.Unlock()
	defer func() {
		t.pendingDialLock.Lock()
		delete(t.pendingDial, random)
		t.pendingDialLock.Unlock()
	}()

	req := &client.Packet{
		Type: client.PacketType_DIAL_REQ,
		Payload: &client.Packet_DialRequest{
			DialRequest: &client.DialRequest{
				Protocol: protocol,
				Address:  address,
				Random:   random,
			},
		},
	}
	klog.V(6).Infof("[tracing] send packet, type: %s", req.Type)

	t.streamLock.Lock()
	err := t.stream.Send(req)
	t.streamLock.Unlock()
	if err != nil {
		return nil, err
	}

	klog.Info("DIAL_REQ sent to proxy server")

	c := &conn{stream: t.stream, tunnel: t}

	select {
	case res := <-resCh:
		if res.err != "" {
			return nil, errors.New(res.err)
		}
		c.connID = res.connid
		c.readCh = make(chan []byte, 10)
		c.closeCh = make(chan string)
		t.connsLock.Lock()
		t.conns[res.connid] = c
		t.connsLock.Unlock()
	case <-time.After(30 * time.Second):
		return nil, errors.New("dial timeout")
	}

	return c, nil
}

func (t *grpcTunnel) Close() {
	close(t.stopCh)
}
