package linked_list

import (
	"errors"
	"fmt"
)

type LinkedList struct {
	Len  int32
	head *node
	tail *node
}

type Entity struct {
	Key   string
	Value string
}

type node struct {
	prev   *node
	next   *node
	entity Entity
}

func (l *LinkedList) TailInsert(entity Entity) (int32, error) {
	newNode := node{
		prev:   l.tail,
		next:   nil,
		entity: entity,
	}
	if l.head == nil {
		l.head = &newNode
	}
	if l.tail != nil {
		l.tail.next = &newNode
	}
	l.tail = &newNode
	l.Len++
	return l.Len, nil
}

func (l *LinkedList) HeadInsert(entity Entity) (int32, error) {
	newNode := node{
		prev:   nil,
		next:   l.head,
		entity: entity,
	}
	if l.head != nil {
		l.head.prev = &newNode
	}
	l.head = &newNode
	if l.tail == nil {
		l.tail = &newNode
	}
	l.Len++
	return l.Len, nil
}

func (l *LinkedList) HeadMove(n *node) (int32, error) {
	n.prev = nil
	n.next = l.head
	l.head.prev = n
	l.head = n
	return l.Len, nil
}

func (l *LinkedList) Remove(n *node) bool {
	if n.prev == nil {
		l.head = n.next
	} else {
		n.prev.next = n.next
	}
	if n.next == nil {
		l.tail = n.prev
	} else {
		n.next.prev = n.prev
	}
	l.Len--
	return true
}

func (l *LinkedList) Search(key string) (Entity, error) {
	n := l.head
	for n != nil {
		if n.entity.Key == key {
			l.HeadMove(n)
			return n.entity, nil
		}
		n = n.next
	}
	return Entity{}, errors.New("key not found")
}

func (l *LinkedList) Export() {
	n := l.head
	for n != nil {
		fmt.Println(fmt.Sprintf("%v", n))
		n = n.next
	}
	return
}

func (l *LinkedList) GetTail() *node {
	return l.tail
}
