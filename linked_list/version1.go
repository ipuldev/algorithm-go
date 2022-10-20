package main

import "fmt"

type Node struct {
	next  *Node
	prev  *Node
	value interface{}
}

type List struct {
	head *Node
	tail *Node
}

func (l *List) append(value interface{}) {
	node := &Node{
		prev:  l.head,
		value: value,
	}

	if l.head != nil {
		node.prev.next = node
	}

	if l.head == nil {
		l.tail = node
	}

	l.head = node
}

func (l *List) Preview() string {
	list := l.head
	str := ""
	for list != nil {
		str += fmt.Sprintf("%v ->", list.value)
		list = list.prev
	}
	return str
}

func (l *List) ReversePreview() string {
	list := l.tail
	str := ""
	for list != nil {
		str += fmt.Sprintf("%v ->", list.value)
		list = list.next
	}
	return str
}

func version1() {
	l := &List{}
	l.append(1)
	l.append(3)
	l.append(5)
	l.append(6)
	fmt.Println("Version 1:")
	fmt.Println("(")
	fmt.Println("\tHead To Tail :")
	resultHead := l.Preview()
	fmt.Println("\t", resultHead)
	fmt.Println("\tTail To Head :")
	resultTail := l.ReversePreview()
	fmt.Println("\t", resultTail)
	fmt.Println(")")
}
