package store

import (
	"sync"
	"maps"
	"slices"
)

type Store[K comparable, V any] struct {
	data map[K]V
	mu   sync.RWMutex
}

var globalStore *Store[string, interface{}]
var once sync.Once

func init() {
	once.Do(func() {
		globalStore = &Store[string, interface{}]{
			data: make(map[string]interface{}),
		}
	})
}

func New[K comparable, V any]() *Store[K, V] {
	return &Store[K, V]{
		data: make(map[K]V),
	}
}

func (s *Store[K, V]) Put(key K, value V) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[key] = value
}

func (s *Store[K, V]) Invalidate(key K) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.data, key)
}

func (s *Store[K, V]) Length() int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return len(s.data)
}

func (s *Store[K, V]) Keys() []K {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return slices.Collect(maps.Keys(s.data))
}

// Helper methods for the global store
func Put(key string, value interface{}) {
	globalStore.Put(key, value)
}

func GetString(key string) string {
	globalStore.mu.Lock()
	defer globalStore.mu.Unlock()

	if v, ok := globalStore.data[key].(string); ok {
		return v
	}
	return ""
}

func GetInt(key string) int {
	globalStore.mu.Lock()
	defer globalStore.mu.Unlock()

	if v, ok := globalStore.data[key].(int); ok {
		return v
	}
	return 0
}

func GetBool(key string) bool {
	globalStore.mu.Lock()
	defer globalStore.mu.Unlock()

	if v, ok := globalStore.data[key].(bool); ok {
		return v
	}
	return false
}

func Clear() {
	globalStore.mu.Lock()
	defer globalStore.mu.Unlock()

	clear(globalStore.data)
}
