package main

import (
	"container/list"
	"fmt"
)

type myLru struct {
	cap int
	m   map[string]*list.Element
	l   *list.List
}

func NewLru(cap int) *myLru {
	return &myLru{
		cap: cap,
		m:   make(map[string]*list.Element),
		l:   list.New(),
	}
}

type entry struct {
	key   string
	value string
}

func (l *myLru) Push(k string, v string) {
	//查看是否存在
	//存在 更改数据，移到最前面
	//不存在 查看是否超出空间长度，没有超出 直接添加到最前面，超出，先将最后面的数据移出，再将新数据添加到最前面
	if data, ok := l.m[k]; ok {
		data.Value = entry{
			key:   k,
			value: v,
		}
		l.m[k] = data
		l.l.MoveToFront(data)
		return
	}

	if len(l.m) >= l.cap {
		e := l.l.Back()
		l.l.Remove(e)
		delete(l.m, e.Value.(entry).key)
	}

	front := l.l.PushFront(entry{
		key:   k,
		value: v,
	})

	l.m[k] = front
}

func (l *myLru) Get(k string) interface{} {
	//查看是否存在
	v, ok := l.m[k]
	if !ok {
		return nil
	}
	//存在  将其移动到最前面 返回数据

	l.l.MoveToFront(v)
	return v.Value.(entry).value
}

func Lru() {
	lru := NewLru(3)
	lru.Push("a", "a")
	lru.Push("b", "b")
	lru.Push("c", "c")
	lru.Push("d", "d")
	fmt.Println(lru.Get("a"))
	fmt.Println(lru.Get("b"))
	fmt.Println(lru.Get("c"))
	fmt.Println(lru.Get("d"))
	lru.Push("e", "e")
	fmt.Println(lru.Get("e"))
}
