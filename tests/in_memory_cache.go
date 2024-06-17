package tests

import "sync"

type InMemoryCache struct {
	cache sync.Map
}

func NewInMemoryCache() *InMemoryCache {
	return &InMemoryCache{
		cache: sync.Map{},
	}
}
func (imc *InMemoryCache) Set(key string, value interface{}) {
	imc.cache.Store(key, value)
}

func (imc *InMemoryCache) Get(key string) (interface{}, bool) {
	return imc.cache.Load(key)
}

func (imc *InMemoryCache) Delete(key string) {
	imc.cache.Delete(key)
}

func (imc *InMemoryCache) Swap(key string, value interface{}) (interface{}, bool) {
	return imc.cache.Swap(key, value)
}

var inMemoryCacheInstance *InMemoryCache
var inMemoryCacheOnce sync.Once

func InMemoryCacheInstance() *InMemoryCache {
	inMemoryCacheOnce.Do(func() {
		inMemoryCacheInstance = NewInMemoryCache()
	})
	return inMemoryCacheInstance
}
