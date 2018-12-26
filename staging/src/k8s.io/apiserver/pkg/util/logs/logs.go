/*
Copyright 2014 The Kubernetes Authors.

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

package logs

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/spf13/pflag"

	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/klog"
)

// logFlushFreq is frequency at which we flush logs after calling InitLogs.
var logFlushFreq time.Duration

// AddFlags registers this package's flags (plus the klog flags used by this package) on arbitrary FlagSets.
// You should call this on *some* flag set.
func AddFlags(fs *pflag.FlagSet) {
	var goflagFS flag.FlagSet
	AddFlagsGoflags(&goflagFS)
	fs.AddGoFlagSet(&goflagFS)
}

// AddFlagsGoflags functions identically to AddFlags, except it works on a standard Go flagset.
// This is to make including these flags in test binaries easier (since those expect standard go flags).
func AddFlagsGoflags(fs *flag.FlagSet) {
	fs.DurationVar(&logFlushFreq, "log-flush-frequency", 5*time.Second, "Maximum number of seconds between log flushes")
	klog.InitFlags(fs)

	// TODO(thockin): This is temporary until we agree on log dirs and put those into each cmd.
	flag.Set("logtostderr", "true")
}

// KlogWriter serves as a bridge between the standard log package and the glog package.
type KlogWriter struct{}

// Write implements the io.Writer interface.
func (writer KlogWriter) Write(data []byte) (n int, err error) {
	klog.InfoDepth(1, string(data))
	return len(data), nil
}

// InitLogs initializes logs the way we want for kubernetes.
func InitLogs() {
	log.SetOutput(KlogWriter{})
	log.SetFlags(0)
	// The default glog flush interval is 5 seconds.
	go wait.Forever(klog.Flush, logFlushFreq)
}

// FlushLogs flushes logs immediately.
func FlushLogs() {
	klog.Flush()
}

// NewLogger creates a new log.Logger which sends logs to klog.Info.
func NewLogger(prefix string) *log.Logger {
	return log.New(KlogWriter{}, prefix, 0)
}

// GlogSetter is a setter to set glog level.
func GlogSetter(val string) (string, error) {
	var level klog.Level
	if err := level.Set(val); err != nil {
		return "", fmt.Errorf("failed set klog.logging.verbosity %s: %v", val, err)
	}
	return fmt.Sprintf("successfully set klog.logging.verbosity to %s", val), nil
}
