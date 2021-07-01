package main

import (
	"fmt"
	"time"
)

type esSort struct {
	Order string
}

var vo = esSort{}

func main() {

	a := make([]int, 0)
	for i := 0; i < 100; i++ {
		go func(i int) {
			a = append(a, i)
		}(i)
	}

	time.Sleep(2*time.Second)

	fmt.Println("slice len ---", len(a))

}
