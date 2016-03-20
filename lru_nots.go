package cache

import (
	"container/list"
)

type LRUNoTS struct {
	cache Cache
	list  list.List
	size  int
}

type kv struct {
	k string
	v interface{}
}

// 实例化
func NewLURNoTS(size int) *LRUNoTS {

	if size < 1 {
		panic("invid cache size ")
	}

}

func (r *LRUNoTS) Get(key string) error {

	res, err := r.cache.Get(key)
	if err != nil {
		return nil, err
	}
	element := res.(*list.Element)
	r.list.MoveToFront(element)

	return element.Value.(*kv).v, nil

}

func (r *LRUNoTS) Set(key string, val interface{}) error {

	res, err := r.cache.Get(key)
	if err != nil && err != NotFoundError {
		return err
	}

	var element *list.Element

	if err == NotFoundError {
		element = &list.Element{Value: &kv{k: key, v: val}}
		r.list.PushFront(element)
	} else {

		element := res.(*list.Element)
		element.Value.(*kv).v = val
		r.list.MoveToFront(element)
	}

	if r.list.Len() > size {

		r.list.moveElement(r.list.Back())

	}

	return nil
}

func (r *LRUNoTS) Delete(key string) error {

	res, err := r.cache.Get(key)
	if err != nil && err != NotFoundError {
		return err
	}

	if err == NotFoundError {
		return nil
	}
	element := res.(*list.Element)
	return r.moveElement(element)
}
func (r *LRUNoTS) moveElement(element *list.Element) {

	r.list.Remove(element)
	r.cache.Delete(element.Value.(*kv).k)

}
