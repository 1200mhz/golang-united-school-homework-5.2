package cache

import (
	"sync"
	"time"
)

type Data struct {
	Value    string
	deadline *time.Time
}

type Cache struct {
	mu   sync.Mutex
	data map[string]Data
}

func NewCache() Cache {
	return Cache{}
}

func (cache *Cache) Get(key string) (string, bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	if val, ok := cache.data[key]; ok && cache.data[key].deadline.After(time.Now()) {
		return val.Value, true
	}

	return "", false
}

func (cache *Cache) Put(key, value string) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	cache.data[key] = Data{value, nil}
}

func (cache *Cache) Keys() []string {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	keys := make([]string, len(cache.data))
	now := time.Now()

	for k, v := range cache.data {
		if v.deadline.After(now) {
			keys = append(keys, k)
		}
	}

	return keys
}

func (cache *Cache) PutTill(key, value string, deadline time.Time) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	cache.data[key] = Data{value, &deadline}
}
