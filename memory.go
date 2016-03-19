package cache

import (
	"fmt"
	"sync"
)

// provide a basic inmemory cache mechanism
type Memory struct {
	sync.Mutex
	cache Cache
}

func NewMemory() Cache {
	return &Memory{
		cache: NewMemoryNoTS(),
	}
}

// provide a key  look a value from cache
func (r *Memory) Get(key string) (interface{}, error) {

	fmt.Println("=============memory")
	r.Lock()
	defer r.Unlock()
	return r.cache.Get(key)
}

// provide a key-value  into the cache
func (r *Memory) Set(key string, value interface{}) error {
	r.Lock()
	defer r.Unlock()
	return r.cache.Set(key, value)
}

// provide a key delete item from the cache
func (r *Memory) Delete(key string) error {
	r.Lock()
	defer r.Unlock()
	return r.cache.Delete(key)

}
