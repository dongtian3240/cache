package cache

//no-thread-safe inmemory
type MemoryNoTS struct {
	items map[string]interface{}
}

func NewMemoryNoTS() *MemoryNoTS {

	return &MemoryNoTS{
		items: map[string]interface{}{},
	}
}
func (r *MemoryNoTS) Get(key string) (interface{}, error) {

	va, ok := r.items[key]
	if !ok {
		return nil, NotFoundError
	}
	return va, nil
}

//
func (r *MemoryNoTS) Set(key string, value interface{}) error {

	r.items[key] = value
	return nil
}

//
func (r *MemoryNoTS) Delete(key string) error {

	delete(r.items, key)
	return nil
}
