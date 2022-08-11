package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
	len  int
}

func (L *LinkedList) InsertAtFirst(data int) {
	node := &Node{
		data: data,
		next: L.head,
	}
	if L.len == 0 {
		L.head = node
		L.len++
		return
	}
	L.head = node
	L.len++

}

func (l *LinkedList) InsertAtEnd(val int) {
	n := Node{}
	n.data = val
	if l.len == 0 {
		l.head = &n
		l.len++
		return
	}
	ptr := l.head
	for i := 0; i < l.len; i++ {
		if ptr.next == nil {
			ptr.next = &n
			l.len++
			return
		}
		ptr = ptr.next
	}
}

func (l *LinkedList) Print() {
	if l.len == 0 {
		fmt.Println("No nodes in list")
	}
	ptr := l.head
	for i := 0; i < l.len; i++ {
		fmt.Println("Node: ", ptr)
		ptr = ptr.next
	}
}

func main() {
	linkedList := LinkedList{}
	linkedList.InsertAtEnd(1)
	//linkedList.InsertAtEnd(2)
	//linkedList.InsertAtEnd(3)
	//linkedList.InsertAtEnd(4)

	linkedList.Print()

}
