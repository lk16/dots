package treesearch

import (
	"sync"

	"github.com/lk16/dots/internal/othello"
)

// CacheKey is identifiying a CacheValue
type CacheKey struct {
	board othello.Board
	depth int
}

// CacheValue is an entry in a Cache associated with a CacheKey
type CacheValue struct {
	alpha int
	beta  int
}

// Cacher allows for saving and looking up heuristic bounds on a board
type Cacher interface {
	Lookup(key CacheKey) (CacheValue, bool)
	Save(key CacheKey, value CacheValue) error
}

// MemoryCache is an in-memory cache
type MemoryCache struct {
	sync.Mutex
	table map[CacheKey]CacheValue
}

var _ Cacher = (*MemoryCache)(nil)

// NewMemoryCache creates a new MemoryCache
func NewMemoryCache() *MemoryCache {
	return &MemoryCache{
		table: make(map[CacheKey]CacheValue, 10000),
	}
}

// Clear empties the cache
func (cacher *MemoryCache) Clear() {
	cacher.Lock()
	defer cacher.Unlock()

	cacher.table = make(map[CacheKey]CacheValue, 10000)
}

// Lookup retrieves heuristic bounds on a board
func (cacher *MemoryCache) Lookup(key CacheKey) (CacheValue, bool) {
	cacher.Lock()
	defer cacher.Unlock()

	entry, ok := cacher.table[key]
	return entry, ok
}

// Save stores heuristic bounds on a board
func (cacher *MemoryCache) Save(key CacheKey, value CacheValue) error {
	cacher.Lock()
	defer cacher.Unlock()

	cacher.table[key] = value
	return nil
}

// NoOpCache creates a no-op memory cache, it doesn't store anything, lookup always fails
type NoOpCache struct{}

var _ Cacher = (*NoOpCache)(nil)

// Lookup retrieves heuristic bounds on a board
func (cacher *NoOpCache) Lookup(key CacheKey) (CacheValue, bool) {
	_ = key
	return CacheValue{}, false
}

// Save stores heuristic bounds on a board
func (cacher *NoOpCache) Save(key CacheKey, value CacheValue) error {
	_ = key
	_ = value
	return nil
}
