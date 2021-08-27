package myheap

import (
	"container/heap"
	"fmt"
	"sort"
	"testing"
)

func TestPriorityQueue(t *testing.T) {

	q := PriorityQueue{{1, "1"}, {5, "5"}, {3, "3"}}

	heap.Init(&q)

	pop := heap.Pop(&q)
	pop = heap.Pop(&q)

	heap.Push(&q, entry{2, "2"})

	pop = heap.Pop(&q)

	fmt.Println(pop)

}

func TestXheap(t *testing.T) {
	h := &HeapInt{2, 1, 5, 4, 8}

	sort.Sort(h)
	fmt.Println(h)

	heap.Init(h)
	heap.Pop(h)
	fmt.Println(h)
}
