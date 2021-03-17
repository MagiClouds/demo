package lru

import "container/list"

type LRUCache struct {
	cap int
	m   map[string]*list.Element
	l   *list.List
}

type entry struct {
	key string
	val string
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		cap: capacity,
		m:   make(map[string]*list.Element, capacity),
		l:   list.New(),
	}
}

func (l *LRUCache) Get(key string) string {
	//查看有没有
	v, ok := l.m[key]
	if !ok {
		//没有直接返回
		return ""
	}

	//存在 将list中的数据放在第一位
	l.l.MoveToFront(v)

	return v.Value.(entry).val
}

func (l *LRUCache) Put(key string, val string) {
	//查看是否存在
	v, ok := l.m[key]
	if ok {
		v.Value = entry{key, val}
		l.l.MoveToFront(v)
		return
	}

	//查看容量是否已满
	n := l.l.Len()

	if n >= l.cap {
		end := l.l.Back()
		l.l.Remove(end)
		delete(l.m, end.Value.(entry).key)
	}

	front := l.l.PushFront(entry{key, val})
	l.m[key] = front
}
