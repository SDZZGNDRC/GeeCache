package geecache

import (
	"GeeCache/lru"
	"errors"
	"sync"
)

type cache struct {
	mu         sync.Mutex
	lru        *lru.Cache
	cacheBytes int64
}

func (c *cache) add(key string, value ByteView) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lru == nil {
		c.lru = lru.New(c.cacheBytes, nil) // Lazy Initialization: Reduce memory take ups, improve performance.
	}
	c.lru.Add(key, value)
}

func (c *cache) get(key string) (value ByteView, err error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lru == nil {
		return ByteView{}, errors.New("KEY NOT EXISTS")
	}

	if v, err := c.lru.Get(key); err == nil {
		return v.(ByteView), nil
	}
	return ByteView{}, errors.New("KEY NOT EXISTS")
}
