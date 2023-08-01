package kvs

import (
	"errors"
	"sync"
)

var (
	ErrNoTransactions = errors.New("there are no transactions")
	ErrStackIsEmpty   = errors.New("tx stack is empty")
)

type Transaction struct {
	store   *sync.Map       // Map to store key-value pairs
	deleted map[string]bool // Map to track deleted keys
	prev    *Transaction    // Parent transaction
}

// Get retrieves a key if it isn't marked as deleted
// It will check its parent transaction if it exists.
func (t *Transaction) Get(key string) (string, error) {
	if t.deleted[key] {
		return "", ErrKeyNotFound
	}

	value, ok := t.store.Load(key)
	if ok {
		return value.(string), nil
	}

	if t.prev != nil {
		return t.prev.Get(key)
	}

	return "", ErrKeyNotFound
}

// Put adds a key-value pair to the store and removes the key from the deleted map if it exists there.
func (t *Transaction) Put(key, value string) error {
	if key == "" || value == "" {
		return ErrKeyValueEmpty
	}

	t.store.Store(key, value)
	delete(t.deleted, key)

	return nil
}

// Delete removes a key-value pair from the store and adds the key to the deleted map.
func (t *Transaction) Delete(key string) error {
	if key == "" {
		return ErrKeyEmpty
	}

	t.store.Delete(key)
	t.deleted[key] = true

	return nil
}

// Commit pushes changes made in the current transaction to the parent transaction.
// If there's no parent, it does nothing.
func (t *Transaction) Commit() {
	if t.prev != nil {
		t.store.Range(func(key, value interface{}) bool {
			t.prev.store.Store(key.(string), value.(string))

			return true
		})

		for key := range t.deleted {
			t.prev.deleted[key] = true
		}
	}
}

type TransactionStack struct {
	top *Transaction // Top transaction in the stack
}

// Push creates a new transaction and pushes it to the top of the stack.
func (s *TransactionStack) Push() {
	newTransaction := &Transaction{
		store:   &sync.Map{},
		deleted: make(map[string]bool),
		prev:    s.top,
	}

	if s.top != nil {
		s.top.store.Range(func(key, value interface{}) bool {
			newTransaction.store.Store(key, value)

			return true
		})
	}

	s.top = newTransaction
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
