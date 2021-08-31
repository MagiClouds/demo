package main

import "fmt"

type mySilce []int

func (s *mySilce) Pop(i int) {
	*s = append((*s)[:i], (*s)[i+1:]...)
}

/**
 *
 * @param num int整型一维数组
 * @param size int整型
 * @return int整型一维数组
 */
func maxInWindows(num []int, size int) []int {
	// write code here
	if len(num) < size || size == 0 {
		return nil
	}

	res := make([]int, 0, len(num)-size+1)
	win := mySilce{}
	for k, v := range num {
		if len(win) > 0 && k-size >= win[0] {
			win.Pop(0)
		}

		for i := 0; i < len(win); {
			if num[win[i]] < v {
				win.Pop(i)
				continue
			}
			i += 1
		}

		if k >= size-1 {
			if len(win) > 0 {
				res = append(res, num[win[0]])
			} else {
				res = append(res, v)
			}
		}

		win = append(win, k)
	}

	return res
}

func main() {
	windows := maxInWindows([]int{2, 3, 4, 2, 6, 2, 5, 1}, 3)
	fmt.Println(windows)
}
