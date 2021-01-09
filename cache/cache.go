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

// Package cache interface 缓存 accessToken 通用接口
package cache

import (
	"errors"
	"time"
)

var (
	// ErrKeyNotFound key不存在
	ErrKeyNotFound = errors.New("key not found")

	// ErrCacheExpired 缓存的值已过期
	ErrCacheExpired = errors.New("cache expired")

	// ErrFlush 清空缓存失败
	ErrFlush = errors.New("unable to flush")

	// ErrSave 存储缓存失败
	ErrSave = errors.New("unable to save")

	// ErrDelete 删除缓存失败
	ErrDelete = errors.New("unable to delete")

	// ErrDecode 缓存解析失败
	ErrDecode = errors.New("unable to decode")
)

// Cache is the top-level cache interface
type Cache interface {

	// Contains 缓存是否存在
	Contains(key string) bool

	// Delete 根据key删除缓存
	Delete(key string) error

	// Fetch 读取一个缓存
	Fetch(key string) (string, error)

	// FetchMulti 读取多个缓存
	FetchMulti(keys []string) map[string]string

	// Flush 清空所有key
	Flush() error

	// Save 保存一个key
	Save(key string, value string, lifeTime time.Duration) error
}
