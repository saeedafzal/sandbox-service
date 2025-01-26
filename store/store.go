package store

import "sync"

type Store struct {
	data map[string]interface{}
	mu   sync.RWMutex
}

var globalStore *Store
var once sync.Once

// Automatically runs when package imported.
func init() {
	once.Do(func() {
		globalStore = &Store{
			data: make(map[string]interface{}),
		}
	})
}

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
