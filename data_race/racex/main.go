package main

import (
	"fmt"
	"sync"
)

func dataRace() int {
	var wg sync.WaitGroup

	c := 0

	for i := 0; i < 10 ; i++  {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j ++ {
				c ++
			}
		}()
	}

	wg.Wait()

	fmt.Println(c)

	return c
}

func main()  {
	dataRace()
}
