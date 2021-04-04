package my_heap

import (
	"container/heap"
	"fmt"
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
