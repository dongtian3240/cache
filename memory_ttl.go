package cache

import (
	"fmt"
	"sync"
	"time"
)

var zeroTTL = time.Duration(0)

type MemoryTTL struct {
	sync.Mutex
	cache      *MemoryNoTS
	setAts     map[string]time.Time
	ttl        time.Duration
	gcInterval time.Duration
}

func NewMemoryWithTTL(ttl time.Duration) *MemoryTTL {

	return &MemoryTTL{
		cache:  NewMemoryNoTS(),
		setAts: map[string]time.Time{},
		ttl:    ttl,
	}
}

func (r *MemoryTTL) StartGC(gcInterval time.Duration) {
	r.gcInterval = gcInterval
	go func() {
		for _ = range time.Tick(r.gcInterval) {
			fmt.Println("gc....")
			for key := range r.cache.items {

				if !r.isValid(key) {
					r.Delete(key)
				}
			}
		}
	}()
}

func (r *MemoryTTL) Delete(key string) error {
	r.Lock()
	defer r.Unlock()
	r.delete(key)
	return nil
}

func (r *MemoryTTL) Get(key string) (interface{}, error) {

	r.Lock()
	defer r.Unlock()
	if !r.isValid(key) {
		r.delete(key)
		return nil, NotFoundError
	}

	value, err := r.cache.Get(key)
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (r *MemoryTTL) Set(key string, value interface{}) error {

	r.Lock()
	defer r.Unlock()
	r.cache.Set(key, value)
	r.setAts[key] = time.Now()
	return nil
}

func (r *MemoryTTL) delete(key string) {
	r.cache.Delete(key)
	delete(r.setAts, key)
}
func (r *MemoryTTL) isValid(key string) bool {

	setAt, ok := r.setAts[key]
	if !ok {
		return false
	}

	if r.ttl == zeroTTL {
		return true
	}

	return setAt.Add(r.ttl).After(time.Now())
}
