package my_heap

type entry struct {
	priority int
	val      string
}

type PriorityQueue []entry

func (pq *PriorityQueue) Len() int {
	return len(*pq)
}

func (pq *PriorityQueue) Less(i, j int) bool {
	return (*pq)[i].priority < (*pq)[j].priority
}

func (pq *PriorityQueue) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}

func (pq *PriorityQueue) Push(data interface{}) {
	*pq = append(*pq, data.(entry))
}

func (pq *PriorityQueue) Pop() interface{} {
	n := len(*pq)
	if n == 0 {
		return nil
	}

	item := (*pq)[n-1]
	*pq = (*pq)[:n-1]

	return item

}
