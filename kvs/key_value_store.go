package kvs

import (
	"errors"
)

var (
	ErrKeyNotFound   = errors.New("key not found")
	ErrKeyValueEmpty = errors.New("key and value both must be non-empty")
	ErrKeyEmpty      = errors.New("key must be non-empty")
)

// KeyValueStore represents the interface for a key-value store.
type KeyValueStore interface {
	Get(key string) (string, error)
	Put(key, value string) error
	Delete(key string) error
}
