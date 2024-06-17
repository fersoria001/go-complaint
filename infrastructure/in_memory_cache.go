package infrastructure

import (
	"go-complaint/infrastructure/queue"
	"log"
	"math"
	"sync"
)

type InMemoryCache struct {
	nextID int
	cache  sync.Map
}

func NewInMemoryCache() *InMemoryCache {
	return &InMemoryCache{
		nextID: 0,
		cache:  sync.Map{},
	}
}
func (imc *InMemoryCache) Set(key string, value interface{}) {
	_, ok := imc.cache.Load(key)
	if !ok {
		newQueue := queue.NewLinkedQueue[interface{}]()
		imc.cache.Store(key, newQueue)
	}
	v, ok := imc.cache.Load(key)
	if !ok {
		log.Println("Error loading cached queue")
		return
	}
	q, ok := v.(*queue.LinkedQueue[interface{}])
	if !ok {
		log.Println("Error casting cached queue")
		return
	}
	q.Enqueue(value)
}

func (imc *InMemoryCache) Get(key interface{}) (interface{}, bool) {
	v, ok := imc.cache.Load(key)
	if !ok {
		return nil, false
	}
	q, ok := v.(*queue.LinkedQueue[interface{}])
	if !ok {
		return nil, false
	}
	d, err := q.Dequeue()
	if err != nil {
		return nil, false
	}
	return d, true
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
