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

	for i := 0; i < 100; i++ {
		go func(i int) {
			vo.Order = fmt.Sprint("i: ", i)
			time.Sleep(time.Millisecond)
			fmt.Println(vo.Order)
		}(i)
	}

	time.Sleep(10*time.Second)

}
