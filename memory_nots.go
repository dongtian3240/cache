package cache

//MemoryNoTS provides a non-thread safe caching mechanism
type MemoryNoTS struct {
	items map[string]interface{}
}

func NewMemoryNoTS() *MemoryNoTS {
	return &MemoryNoTS{
		items: map[string]interface{}{},
	}
}

// provide a key  look a value from cache
func (r *MemoryNoTS) Get(key string) (interface{}, error) {

	va, ok := r.items[key]
	if !ok {
		return nil, NotFoundError
	}
	return va, nil

}

// provide a key-value  into the cache
func (r *MemoryNoTS) Set(key string, value interface{}) error {

	r.items[key] = value
	return nil
}

// provide a key delete item from the cache
func (r *MemoryNoTS) Delete(key string) error {

	delete(r.items, key)

	return nil
}
