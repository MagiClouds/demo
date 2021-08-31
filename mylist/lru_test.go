package main

import (
	"fmt"
	"testing"
)



func TestLru(t *testing.T) {
	lru := NewLru(3)
	lru.Push("a", "a")
	lru.Push("b", "b")
	lru.Push("c", "c")
	lru.Push("d", "d")
	fmt.Println(lru.Get("a"))
	fmt.Println(lru.Get("b"))
	fmt.Println(lru.Get("c"))
	fmt.Println(lru.Get("d"))
	lru.Push("e", "e")
	fmt.Println(lru.Get("b"))
	fmt.Println(lru.Get("c"))
	lru.Push("f", "f")

	fmt.Println(lru.Get("c"))
	fmt.Println(lru.Get("d"))


}
