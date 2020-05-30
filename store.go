package cache

import (
	"sync"
)

// Store holds items in memory.
type Store struct {
	mu    sync.RWMutex
	items map[string]*Item
}

// Delete an item from the store.
func (s *Store) Delete(key string) {
	s.mu.Lock()
	delete(s.items, key)
	s.mu.Unlock()
}

// Add add an item to the store, the item
// value can be type of interface{}.
func (s *Store) Add(key string, value interface{}) {
	s.mu.Lock()
	s.items[key] = &Item{
		data: value,
	}
	s.mu.Unlock()
}

// Get an item from the storem, returns
// a pointer to item.
func (s *Store) Get(key string) (i *Item) {
	s.mu.RLock()
	i = s.items[key]
	s.mu.RUnlock()
	return i
}
