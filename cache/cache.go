// Package cache 缓存 accessToken 通用接口
package cache

import (
	"errors"
	"time"
)

var (
	// ErrKeyNotFound key不存在
	ErrKeyNotFound = errors.New("key not found")
	// ErrCacheExpired returns an error when the cache key was expired.
	ErrCacheExpired = errors.New("cache expired")

	// ErrFlush returns an error when flush fails.
	ErrFlush = errors.New("unable to flush")

	// ErrSave returns an error when save fails.
	ErrSave = errors.New("unable to save")

	// ErrDelete returns an error when deletion fails.
	ErrDelete = errors.New("unable to delete")

	// ErrDecode returns an errors when decode fails.
	ErrDecode = errors.New("unable to decode")
)

// Cache is the top-level cache interface
type Cache interface {

	// Contains check if a cached key exists
	Contains(key string) bool

	// Delete remove the cached key
	Delete(key string) error

	// Fetch retrieve the cached key value
	Fetch(key string) (string, error)

	// FetchMulti retrieve multiple cached keys value
	FetchMulti(keys []string) map[string]string

	// Flush remove all cached keys
	Flush() error

	// Save cache a value by key
	Save(key string, value string, lifeTime time.Duration) error
}
