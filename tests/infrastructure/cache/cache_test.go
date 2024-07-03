package cache_test

import (
	"go-complaint/infrastructure/cache"
	"testing"
)

func TestCache(t *testing.T) {
	go cache.Cache(cache.RequestChannel)
	cache.SendToChannel(cache.RequestChannel, cache.Request{
		Type:    cache.WRITE,
		Key:     "test",
		Payload: "test",
	})
	out := make(chan cache.Request, 1)
	cache.SendToChannel(cache.RequestChannel, cache.Request{
		Type: cache.READ,
		Key:  "test",
		Out:  out,
	})
	req := <-out
	t.Log(req)
}
