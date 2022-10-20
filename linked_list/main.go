package main

import "log"

type Node struct {
	next  *Node
	prev  *Node
	value interface{}
}

type List struct {
	head *Node
	tail *Node
}

func (l *List) append(data interface{}) {
	current := l.head
	if l.head == nil {
		current = &Node{
			value: data,
		}
		return
	}

	for current.next != nil {
		current = current.next
	}
	current.next = &Node{
		value: data,
	}
}

func main() {
	l := &List{}
	l.append(1)
	log.Println(l)
}
