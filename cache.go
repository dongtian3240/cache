package cache

type Cache interface {

	// provide a key  look a value from cache
	Get(key string) (interface{}, error)

	// provide a key-value  into the cache
	Set(key string, value interface{}) error

	// provide a key delete item from the cache
	Delete(key string) error
}
