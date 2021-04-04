package main

import (
	"container/list"
	"fmt"
)

func main() {

	l := list.New()

	e1 := l.PushBack(1)

	l.PushBack(3)
	l.PushBack(4)
	e5 := l.PushBack(5)
	e6 := l.PushBack(6)

	fmt.Printf("e1： %+v \n", e1)
	fmt.Printf("e5： %+v \n", e5)
	fmt.Printf("e6: %+v \n", e6)

	fmt.Printf("%+v \n", l)

	end := l.Back()

	fmt.Printf("%+v \n", end)

	start := l.Front()

	fmt.Printf("%+v \n", start)

	l.MoveToFront(e5)

	fmt.Println(l.Front())

	e7 := l.PushBack(8)
	fmt.Println(e7)

	l.Remove(e5)
	fmt.Println(l.Front())




}
