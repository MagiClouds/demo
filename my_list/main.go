package main

import (
	"container/list"
	"fmt"
)

func main()  {

	l := list.New()

	back := l.PushBack(1)

	fmt.Printf("%+v \n", back)

	fmt.Printf("%+v \n", l)

	element := l.Back()

	fmt.Printf("%+v", element)

}
