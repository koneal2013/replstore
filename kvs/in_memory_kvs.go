package kvs

// InMemoryKVStore is an in-memory implementation of KeyValueStore.
type InMemoryKVStore struct {
	data map[string]string
}

func NewInMemoryKVStore() *InMemoryKVStore {
	return &InMemoryKVStore{
		data: make(map[string]string),
	}
}

// Get retrieves the value for the given key.
func (s *InMemoryKVStore) Get(key string) (string, error) {
	value, ok := s.data[key]
	if !ok {
		return "", ErrKeyNotFound
	}

	return value, nil
}

// Put stores the value for the given key.
func (s *InMemoryKVStore) Put(key, value string) error {
	if key == "" || value == "" {
		return ErrKeyValueEmpty
	}

	s.data[key] = value

	return nil
}

// Delete removes the key-value pair for the given key.
func (s *InMemoryKVStore) Delete(key string) error {
	delete(s.data, key)

	return nil
}
