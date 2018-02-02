/*
Copyright 2017 The Kubernetes Authors.

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

package azure

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	fakeCacheTTL = 2 * time.Second
)

type fakeDataObj struct{}

type fakeDataSource struct {
	data map[string]*fakeDataObj
	lock sync.Mutex
}

func (fake *fakeDataSource) get(key string) (interface{}, error) {
	fake.lock.Lock()
	defer fake.lock.Unlock()

	if v, ok := fake.data[key]; ok {
		return v, nil
	}

	return nil, nil
}

func (fake *fakeDataSource) list() (map[string]interface{}, error) {
	fake.lock.Lock()
	defer fake.lock.Unlock()

	result := make(map[string]interface{})
	for k, v := range fake.data {
		result[k] = v
	}
	return result, nil
}

func (fake *fakeDataSource) set(data map[string]*fakeDataObj) {
	fake.lock.Lock()
	defer fake.lock.Unlock()

	fake.data = data
}

func newFakeCache() (*fakeDataSource, *timedCache) {
	dataSource := &fakeDataSource{
		data: make(map[string]*fakeDataObj),
	}
	getter := dataSource.get
	lister := dataSource.list
	return dataSource, newTimedcache(fakeCacheTTL, getter, lister)
}

func TestCacheGet(t *testing.T) {
	val := &fakeDataObj{}
	cases := []struct {
		name     string
		data     map[string]*fakeDataObj
		key      string
		expected interface{}
	}{
		{
			name:     "cache should return nil for empty data source",
			key:      "key1",
			expected: nil,
		},
		{
			name:     "cache should return nil for non exist key",
			data:     map[string]*fakeDataObj{"key2": val},
			key:      "key1",
			expected: nil,
		},
		{
			name:     "cache should return data for existing key",
			data:     map[string]*fakeDataObj{"key1": val},
			key:      "key1",
			expected: val,
		},
	}

	for _, c := range cases {
		dataSource, cache := newFakeCache()
		dataSource.set(c.data)
		val, err := cache.Get(c.key)
		assert.NoError(t, err, c.name)
		assert.Equal(t, c.expected, val, c.name)
	}
}

func TestCacheList(t *testing.T) {
	val := &fakeDataObj{}
	cases := []struct {
		name     string
		data     map[string]*fakeDataObj
		expected []interface{}
	}{
		{
			name:     "cache should get empty result with empty data source",
			expected: []interface{}{},
		},
		{
			name:     "cache should get same data with provided data source",
			data:     map[string]*fakeDataObj{"key1": val},
			expected: []interface{}{val},
		},
	}

	for _, c := range cases {
		dataSource, cache := newFakeCache()
		dataSource.set(c.data)
		val, err := cache.List()
		assert.NoError(t, err, c.name)
		assert.Equal(t, c.expected, val, c.name)
	}
}

func TestCacheExpired(t *testing.T) {
	key := "key1"
	val := &fakeDataObj{}
	data := map[string]*fakeDataObj{
		key: val,
	}
	dataSource, cache := newFakeCache()
	dataSource.set(data)

	v, err := cache.Get(key)
	assert.NoError(t, err)
	assert.Equal(t, val, v, "cache should get correct data")

	cache.Delete(key)
	v, err = cache.Get(key)
	assert.NoError(t, err)
	assert.Equal(t, nil, v, "cache should get nil after data is removed")
	valList, err := cache.List()
	assert.NoError(t, err)
	assert.Equal(t, 0, len(valList), "cache should get empty list after data is removed")

	time.Sleep(fakeCacheTTL)
	v, err = cache.Get(key)
	assert.NoError(t, err)
	assert.Equal(t, val, v, "cache should get correct data even after expired")
}
