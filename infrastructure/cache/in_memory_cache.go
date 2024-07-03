package cache

import (
	"math"
	"sync"
)

type PublisherMap = map[string]interface{}

const PUBLISHER_MAP = "publisherMap"

type InMemoryCache struct {
	nextID int
	cache  sync.Map
}

func NewInMemoryCache() *InMemoryCache {
	imc := &InMemoryCache{
		nextID: 0,
		cache:  sync.Map{},
	}
	var publisherMap PublisherMap = map[string]interface{}{}
	imc.cache.Store(PUBLISHER_MAP, publisherMap)
	return imc
}

func (imc *InMemoryCache) GetPublisherMap() PublisherMap {
	v, ok := imc.cache.Load(PUBLISHER_MAP)
	if !ok {
		return nil
	}
	return v.(PublisherMap)
}

func (imc *InMemoryCache) SetPublish(id string, value interface{}) {
	publisherMap := imc.GetPublisherMap()
	publisherMap[id] = value
	imc.cache.Swap(PUBLISHER_MAP, publisherMap)
}

func (imc *InMemoryCache) GetPublish(id string) (interface{}, bool) {
	publisherMap := imc.GetPublisherMap()
	v, ok := publisherMap[id]
	if !ok {
		return nil, false
	}
	return v, true
}

func (imc *InMemoryCache) ClearPublisherMap() {
	var publisherMap PublisherMap = map[string]interface{}{}
	imc.cache.Swap(PUBLISHER_MAP, publisherMap)
}

func (imc *InMemoryCache) Set(key string, value interface{}) {
	imc.cache.Store(key, value)
}

func (imc *InMemoryCache) Get(key interface{}) (interface{}, bool) {
	v, ok := imc.cache.Load(key)
	if !ok {
		return nil, false
	}
	return v, true
}

func (imc *InMemoryCache) Delete(key string) {
	imc.cache.Delete(key)
}

func (imc *InMemoryCache) Swap(key string, value interface{}) (interface{}, bool) {
	return imc.cache.Swap(key, value)
}

func (imc *InMemoryCache) NextID() int {
	if imc.nextID > math.MaxInt32 {
		imc.nextID = 0
	}
	imc.nextID++
	return imc.nextID
}

var inMemoryCacheInstance *InMemoryCache
var inMemoryCacheOnce sync.Once

func InMemoryCacheInstance() *InMemoryCache {
	inMemoryCacheOnce.Do(func() {
		inMemoryCacheInstance = NewInMemoryCache()
	})
	return inMemoryCacheInstance
}
