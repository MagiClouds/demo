package slice2stack

import (
	"fmt"
	"testing"
)

func TestNewStack(t *testing.T) {
	stack := NewStack()
	fmt.Println(stack.Pop())
	stack.Put(1)
	stack.Put(2)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	stack.Put(3)
	fmt.Println(stack.Pop())
}