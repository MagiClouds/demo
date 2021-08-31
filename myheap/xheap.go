package myheap

type HeapInt []int

func (h HeapInt) Len() int {
	return len(h)
}

func (h HeapInt) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h HeapInt) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *HeapInt) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *HeapInt) Pop() interface{} {
	if len(*h) == 0 {
		return nil
	}
	data := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return data
}

