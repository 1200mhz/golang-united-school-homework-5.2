package cache

import (
	"sync"
	"time"
)

type Data struct {
	value    string
	deadline *time.Time
}

type Cache struct {
	mu   sync.Mutex
	data map[string]Data
}

func NewCache() Cache {
	return Cache{
		data: map[string]Data{},
	}
}

func (cache *Cache) Get(key string) (string, bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	if _, ok := cache.data[key]; !ok {
		return "", false
	}

	if cache.data[key].deadline != nil && cache.data[key].deadline.Before(time.Now()) {
		return "", false
	}

	return cache.data[key].value, true
}

func (cache *Cache) Put(key, value string) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	cache.data[key] = Data{value, nil}
}

func (cache *Cache) Keys() []string {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	var keys []string
	now := time.Now()

	for k, v := range cache.data {
		if v.deadline != nil && v.deadline.Before(now) {
			continue
		}

		keys = append(keys, k)
	}

	return keys
}

func (cache *Cache) PutTill(key, value string, deadline time.Time) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	cache.data[key] = Data{value, &deadline}
}
