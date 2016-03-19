package cache

import (
	"fmt"
	"sync"
	"time"
)

var zeroTTL = time.Duration(0)

type MemoryExpire struct {
	sync.Mutex
	cache      *MemoryNoTS
	setAts     map[string]time.Time
	expire     time.Duration
	gcInterval time.Duration
}

func NewMemoryWithExpire(expire time.Duration) *MemoryExpire {

	return &MemoryExpire{
		cache:  NewMemoryNoTS(),
		setAts: map[string]time.Time{},
		expire: expire,
	}
}

func (r *MemoryExpire) StartGC(gcInterval time.Duration) {
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

func (r *MemoryExpire) Delete(key string) error {
	r.Lock()
	defer r.Unlock()
	r.delete(key)
	return nil
}

func (r *MemoryExpire) Get(key string) (interface{}, error) {

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

func (r *MemoryExpire) Set(key string, value interface{}) error {

	r.Lock()
	defer r.Unlock()
	r.cache.Set(key, value)
	r.setAts[key] = time.Now().Add(r.expire)
	return nil
}

func (r *MemoryExpire) delete(key string) {
	r.cache.Delete(key)
	delete(r.setAts, key)
}
func (r *MemoryExpire) isValid(key string) bool {

	setAt, ok := r.setAts[key]
	if !ok {
		return false
	}

	if r.expire == zeroTTL {
		return true
	}
	b := time.Now().Before(setAt)
	fmt.Println(b)
	return b
}
