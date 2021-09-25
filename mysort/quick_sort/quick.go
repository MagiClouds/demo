package quick_sort

func QuickSort(l []int) []int {
	if len(l) <= 1 {
		return l
	}

	temp := l[0]
	l1 := make([]int, 0, len(l))
	l2 := make([]int, 0, len(l))

	for _, v := range l[1:] {
		if v <= temp {
			l1 = append(l1, v)
		}else {
			l2 = append(l2, v)
		}
	}

	res := append(QuickSort(l1), temp)
	return  append(res, QuickSort(l2)...)
}
