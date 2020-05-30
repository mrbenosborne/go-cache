package cache

import "sync"

// Cache holds many cache stores.
type Cache struct {
	mu     sync.RWMutex
	Stores map[string]*Store
}

// New create a new Cache.
func New(size int) *Cache {
	return &Cache{
		Stores: make(map[string]*Store, size),
	}
}

// New make a new store, if an existing store
// exists with the same name then it will be
// overwritten.
//
// Returns a pointer to type of Store.
func (c *Cache) New(name string) *Store {
	s := &Store{
		items: make(map[string]*Item),
	}
	c.mu.Lock()
	c.Stores[name] = s
	c.mu.Unlock()
	return s
}

// Get a store by name, returns nil if the store
// does not exist.
func (c *Cache) Get(name string) *Store {
	var s *Store
	c.mu.RLock()
	s = c.Stores[name]
	c.mu.RUnlock()
	return s
}

// Delete a store.
func (c *Cache) Delete(name string) {
	c.mu.Lock()
	delete(c.Stores, name)
	c.mu.Unlock()
}
