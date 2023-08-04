package kvs

import (
	"errors"
)

var (
	ErrNoTransactions = errors.New("there are no transactions")
	ErrStackIsEmpty   = errors.New("tx stack is empty")
)

type Transaction struct {
	store   map[string]string // Map to store key-value pairs
	deleted map[string]bool   // Map to track deleted keys
	prev    *Transaction      // Parent transaction
	kvStore KeyValueStore
}

func NewTransaction(kvStore KeyValueStore) *Transaction {
	return &Transaction{
		store:   make(map[string]string),
		deleted: make(map[string]bool),
		kvStore: kvStore,
	}
}

// Get retrieves a key if it isn't marked as deleted
// It will check its parent transaction if it exists.
func (t *Transaction) Get(key string) (string, error) {
	if t.deleted[key] {
		return "", ErrKeyNotFound
	}

	value, ok := t.store[key]
	if ok {
		return value, nil
	}

	if t.prev != nil {
		return t.prev.Get(key)
	}

	return t.kvStore.Get(key)
}

// Put adds a key-value pair to the store and removes the key from the deleted map if it exists there.
func (t *Transaction) Put(key, value string) error {
	if key == "" || value == "" {
		return ErrKeyValueEmpty
	}

	t.store[key] = value
	delete(t.deleted, key)

	return nil
}

// Delete removes a key-value pair from the store and adds the key to the deleted map.
func (t *Transaction) Delete(key string) error {
	if key == "" {
		return ErrKeyEmpty
	}

	delete(t.store, key)
	t.deleted[key] = true

	return nil
}

// Commit pushes changes made in the current transaction to the parent transaction.
func (t *Transaction) Commit() {
	if t.prev != nil {
		for key, value := range t.store {
			t.prev.store[key] = value
		}

		for key := range t.deleted {
			t.prev.deleted[key] = true
		}
	} else if t.kvStore != nil {
		for key, value := range t.store {
			_ = t.kvStore.Put(key, value)
		}
	}
}

type TransactionStack struct {
	top *Transaction // Top transaction in the stack
}

// Push creates a new transaction and pushes it to the top of the stack.
func (s *TransactionStack) Push(newTx *Transaction) {
	newTx.prev = s.top

	if s.top != nil {
		for key, value := range s.top.store {
			newTx.store[key] = value
		}
	}

	s.top = newTx
}

// Pop pops the top transaction from the stack.
func (s *TransactionStack) Pop() error {
	if s.top == nil {
		return ErrStackIsEmpty
	}

	if s.top != nil {
		s.top = s.top.prev
	}

	return nil
}

// Current returns the transaction at the top of the stack.
func (s *TransactionStack) Current() (*Transaction, error) {
	if s.top == nil {
		return nil, ErrNoTransactions
	}

	return s.top, nil
}
