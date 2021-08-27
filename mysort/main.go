package main

import (
	"fmt"
	"sort"
)

type People struct {
	Name string
	Age  int32
}

type PeopleList []People

func (p PeopleList) Len() int {
	return len(p)
}

func (p PeopleList) Less(i, j int) bool {
	return p[i].Age < p[j].Age
}

func (p PeopleList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {

	p := PeopleList{{Age: 2}, {Age: 5}, {Age: 3}}

	sort.Sort(p)

	fmt.Println(p)

}
