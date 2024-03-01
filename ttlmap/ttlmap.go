package ttlmap

import (
	"github.com/patrickmn/go-cache"

	"sync"
	"time"
)

type ConcurrentMap struct {
	m     *cache.Cache
	mutex sync.RWMutex
}

func NewConcurrentMap(defaultExpiration time.Duration, cleanupInterval time.Duration) *ConcurrentMap {
	return &ConcurrentMap{
		m: cache.New(defaultExpiration, cleanupInterval),
	}
}

func (c *ConcurrentMap) Set(key string, value interface{}, duration time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.m.Set(key, value, duration)
}

func (c *ConcurrentMap) Get(key string) (interface{}, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.m.Get(key)
}

func (c *ConcurrentMap) Delete(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.m.Delete(key)
}
