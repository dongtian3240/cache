package cache

import (
	"fmt"
	"testing"
	"time"
)

func TestTTL(t *testing.T) {

	cache := NewMemoryWithTTL(2 * time.Second)
	cache.StartGC(time.Millisecond * 10)
	cache.Set("test_key", "test_data")
	time.Sleep(200 * time.Millisecond)
	v, err := cache.Get("test_key")
	if err == nil {
		t.Fatal("data found")
	}

	fmt.Println("value = ", v)

}
