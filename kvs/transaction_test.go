package kvs

import (
	"errors"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransaction_Put(t *testing.T) {
	trans := Transaction{store: &sync.Map{}, deleted: make(map[string]bool)}
	err := trans.Put("test", "value")
	require.NoError(t, err)

	val, err := trans.Get("test")
	require.NoError(t, err)

	require.Equal(t, "value", val)
}

func TestTransaction_Delete(t *testing.T) {
	trans := Transaction{store: &sync.Map{}, deleted: make(map[string]bool)}
	err := trans.Put("test", "value")
	require.NoError(t, err)

	err = trans.Delete("test")
	require.NoError(t, err)

	val, err := trans.Get("test")
	if !errors.Is(err, ErrKeyNotFound) {
		t.Errorf("Expected %v error, got %v", ErrKeyNotFound, err)
	}

	require.Error(t, err)
	require.Equal(t, ErrKeyNotFound, err)
	require.Empty(t, val)
}

func TestTransactionStack(t *testing.T) {
	stack := TransactionStack{}

	stack.Push()

	require.NotNil(t, stack.Current(), "Expected current transaction, got nil")

	trans := stack.Current()
	err := trans.Put("test", "value")
	require.NoError(t, err)

	stack.Pop()

	require.Nil(t, stack.Current(), "Expected nil current transaction, got non-nil")
}
