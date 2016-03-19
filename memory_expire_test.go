package cache

import (
	"fmt"
	"testing"
	"time"
)

func TestExpire(t *testing.T) {

	cache := NewMemoryWithExpire(2 * time.Second)
	cache.StartGC(time.Millisecond * 100)
	cache.Set("test_key", "test_data")
	time.Sleep(5 * time.Second)
	v, err := cache.Get("test_key")
	if err == nil {
		t.Fatal("data found")
	}

	fmt.Println("value = ", v)

}
