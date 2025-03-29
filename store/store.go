package store

import "sync"

type Store[K comparable, V any] struct {
	data map[K]V
	mu   sync.RWMutex
}

var globalStore *Store[string, interface{}]
var once sync.Once

// Automatically runs when package imported.
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

// Helper methods for the global store
func Put(key string, value interface{}) {
	globalStore.mu.Lock()
	defer globalStore.mu.Unlock()
	globalStore.data[key] = value
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
