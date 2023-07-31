package kvs

// KeyValueStore represents the interface for a key-value store.
type KeyValueStore interface {
	Get(key string) (string, error)
	Put(key, value string) error
	Delete(key string) error
}
