// MIT License

// Copyright (c) 2019 Berryhe

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Package cache 缓存接口的默认内存实现
package cache

import (
	"sync"
	"time"
)

type (
	syncMapItem struct {
		data     string
		duration int64
	}

	syncMap struct {
		storage *sync.Map
	}
)

// NewDefaultCache 默认使用内存缓存
func NewDefaultCache() Cache {
	return &syncMap{&sync.Map{}}
}

func (sm *syncMap) read(key string) (*syncMapItem, error) {
	v, ok := sm.storage.Load(key)
	if !ok {
		return nil, ErrKeyNotFound
	}

	item := v.(*syncMapItem)

	if item.duration == 0 {
		return item, nil
	}

	if item.duration <= time.Now().Unix() {
		_ = sm.Delete(key)
		return nil, ErrCacheExpired
	}

	return item, nil
}

// Contains key是否存在
func (sm *syncMap) Contains(key string) bool {
	_, err := sm.Fetch(key)
	return err == nil
}

// Delete 将key删除
func (sm *syncMap) Delete(key string) error {
	sm.storage.Delete(key)
	return nil
}

// Fetch 读取一个key
func (sm *syncMap) Fetch(key string) (string, error) {
	item, err := sm.read(key)
	if err != nil {
		return "", err
	}

	return item.data, nil
}

// FetchMulti 读取多个key
func (sm *syncMap) FetchMulti(keys []string) map[string]string {
	result := make(map[string]string)

	for _, key := range keys {
		if value, err := sm.Fetch(key); err == nil {
			result[key] = value
		}
	}

	return result
}

// Flush 删除所有key
func (sm *syncMap) Flush() error {
	sm.storage = &sync.Map{}
	return nil
}

// Save 存一个键值到syncMap
func (sm *syncMap) Save(key string, value string, lifeTime time.Duration) error {
	var duration int64

	if lifeTime > 0 {
		duration = time.Now().Unix() + int64(lifeTime.Seconds())
	}

	sm.storage.Store(key, &syncMapItem{value, duration})
	return nil
}
