package myheap

import "container/heap"

type PriorityQueue []int

func (p *PriorityQueue) Len() int {
	return len(*p)
}

func (p *PriorityQueue) Less (i, j int) bool {
	return (*p)[i] > (*p)[j]
}

func (p *PriorityQueue) Swap(i, j int) {
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}

func (p *PriorityQueue) Pop() interface{} {
	if len(*p) == 0 {
		return nil
	}
	data := (*p)[len(*p)-1]
	*p = (*p)[:len(*p)-1]
	return data
}

func (p *PriorityQueue) Push(data interface{}) {
	*p = append(*p, data.(int))
}


/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param input int整型一维数组
 * @param k int整型
 * @return int整型一维数组
 */
func GetLeastNumbers_Solution( input []int ,  k int ) []int {
	// write code here
	if len(input) <= k {
		return input
	}

	res := &PriorityQueue{}
	heap.Init(res)

	for _, v := range input {
		if len(*res) < k {
			heap.Push(res, v)
			continue
		}
		pop := heap.Pop(res)
		if pop.(int) > v {
			heap.Push(res, v)
		}else {
			heap.Push(res, pop.(int))
		}
	}

	bo := make([]int, 0, len(*res))
	for _, v := range *res {
		bo = append(bo, v)
	}

	return bo
}