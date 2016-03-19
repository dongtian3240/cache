package cache

import (
	"fmt"
	"testing"
)

func TestMemoryGet(t *testing.T) {
	cache := NewMemory()
	cache.Set("name", "冬天")
	name, err := cache.Get("name")
	fmt.Println("name=", name, "key=name", "error = ", err)
	//t.Log()
}
