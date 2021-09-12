package main

import (
"fmt"
	"time"
)

func main() {
	f()
	fmt.Println("Returned normally from f.")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
			fmt.Println("Recovered in f", r)
		}
	}()

	fmt.Println("Calling g.")

	go func() {
		g()
	}()

	time.Sleep(1*time.Second)

	fmt.Println("Returned normally from g.")
}

func g() {
	k()
}

func k() {
	a := []string{}
	fmt.Println(a[1])
}
