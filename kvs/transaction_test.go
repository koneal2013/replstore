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

func TestTransaction_Revert(t *testing.T) {
	// Create parent transaction
	parent := &Transaction{
		store:   &sync.Map{},
		deleted: map[string]bool{},
		prev:    nil,
	}

	parent.store.Store("keyExistsInParent", true)
	parent.deleted["keyDeletedInParent"] = true

	// Create current transaction with a link to its parent
	trans := &Transaction{
		store:   &sync.Map{},
		deleted: map[string]bool{},
		prev:    parent,
	}

	trans.store.Store("keyExistsInChild", true)
	trans.deleted["keyDeletedInChild"] = true

	// Revert transaction
	trans.Revert()

	// Check if keys are correctly reverted
	_, exists := parent.store.Load("keyExistsInParent")
	require.True(t, exists, "Key 'keyExistsInParent' was not correctly reverted in store")

	_, exists = parent.deleted["keyDeletedInParent"]
	require.True(t, exists, "Key 'keyDeletedInParent' was not correctly reverted in deleted")

	_, exists = parent.store.Load("keyExistsInChild")
	require.False(t, exists, "Key 'keyExistsInChild' was not correctly reverted in store")

	_, exists = parent.deleted["keyDeletedInChild"]
	require.False(t, exists, "Key 'keyDeletedInChild' was not correctly reverted in deleted")
}
