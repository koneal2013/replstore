package kvs

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInMemoryKeyValueStore(t *testing.T) {
	store := &InMemoryKeyValueStore{}

	// Test Put and Get
	err := store.Put("key1", "value1")
	require.NoError(t, err)

	value, err := store.Get("key1")
	require.NoError(t, err)
	require.Equal(t, "value1", value)

	// Test Delete
	err = store.Delete("key1")
	require.NoError(t, err)

	_, err = store.Get("key1")
	require.Error(t, err)
	require.Equal(t, err, ErrKeyNotFound)
}
