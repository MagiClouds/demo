package lru

import (
	"fmt"
	"testing"
)

func TestLRUCache(t *testing.T) {

	cache := NewLRUCache(3)

	cache.Put("1", "1")
	cache.Put("2", "2")
	cache.Put("3", "3")

	cache.Put("4", "4")
	cache.Put("5", "5")
	cache.Put("6", "6")


	get1 := cache.Get("4")

	fmt.Println(get1)
}
