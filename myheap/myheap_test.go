package myheap

import (
	"container/heap"
	"fmt"
	"sort"
	"testing"
)

//func TestPriorityQueue(t *testing.T) {
//
//	q := PriorityQueue{{1, "1"}, {5, "5"}, {3, "3"}}
//
//	heap.Init(&q)
//
//	pop := heap.Pop(&q)
//	pop = heap.Pop(&q)
//
//	heap.Push(&q, entry{2, "2"})
//
//	pop = heap.Pop(&q)
//
//	fmt.Println(pop)
//
//}

func TestGetLeastNumbers_Solution(t *testing.T)  {
	fmt.Println(GetLeastNumbers_Solution([]int{4,5,1,6,2,7,3,8}, 4))
}

func TestXheap(t *testing.T) {
	h := &HeapInt{2, 1, 5, 4, 8}

	sort.Sort(h)
	fmt.Println(h)

	heap.Init(h)
	heap.Pop(h)
	fmt.Println(h)
}

func TestSlice(t *testing.T) {
	a := make([]int, 0)
	fmt.Println(a)
	fmt.Println(len(a))

	a = nil
	fmt.Println(a)
	fmt.Println(len(a))
}
