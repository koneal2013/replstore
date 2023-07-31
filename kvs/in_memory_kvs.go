package kvs

import (
	"sync"
)

// InMemoryKeyValueStore is an in-memory implementation of KeyValueStore.
type InMemoryKeyValueStore struct {
	data sync.Map
}

// Get retrieves the value for the given key.
func (s *InMemoryKeyValueStore) Get(key string) (string, error) {
	value, ok := s.data.Load(key)
	if !ok {
		return "", ErrKeyNotFound
	}

	return value.(string), nil
}

// Put stores the value for the given key.
func (s *InMemoryKeyValueStore) Put(key, value string) error {
	s.data.Store(key, value)

	return nil
}

// Delete removes the key-value pair for the given key.
func (s *InMemoryKeyValueStore) Delete(key string) error {
	s.data.Delete(key)

	return nil
}
