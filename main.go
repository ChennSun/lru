package main

import (
	"fmt"
	"lru/linked_list"
	"sync"
)

type Lru struct {
	lock  sync.RWMutex
	list  *linked_list.LinkedList
	limit int32
}

var lru = Lru{
	list:  &linked_list.LinkedList{},
	limit: 10,
}

func add(key, value string) {
	lru.lock.Lock()
	defer lru.lock.Unlock()
	if lru.list.Len == lru.limit {
		lru.list.Remove(lru.list.GetTail())
	}
	lru.list.HeadInsert(linked_list.Entity{
		Key:   key,
		Value: value,
	})
}

func get(key string) string {
	e, err := lru.list.Search(key)
	if err != nil {
		fmt.Println(err)
	}
	return e.Value
}

func main() {
	for i := 0; i < 20; i++ {
		add(fmt.Sprintf("%d", i), fmt.Sprintf("value-%d", i))
		lru.list.Export()
		fmt.Println("-----")
	}
}
