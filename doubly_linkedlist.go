package main

import "fmt"

type Node struct {
	Data interface{}

	Next *Node
	Prev *Node
}

type DoublyLinkedList struct {
	head *Node
	len  int
}

func (l *DoublyLinkedList) Insert(data interface{}) {
	node := &Node{
		Data: data,
		Next: l.head,
	}
	if l.head != nil {
		l.head.Prev = node
	}
	l.head = node
	l.len++
}

func (l *DoublyLinkedList) Print() {
	ptr := l.head
	for i := 0; i < l.len; i++ {
		fmt.Println("Node: ", ptr)
		ptr = ptr.Next
	}
}

func main() {
	linkedList := DoublyLinkedList{}
	linkedList.Insert(1)
	linkedList.Insert(2)
	linkedList.Insert(3)
	linkedList.Insert(4)

	linkedList.Print()

}
