package data_structures

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) Append(val int) {
	new_node := &Node{data: val}

	if l.head == nil {
		l.head = new_node
		return
	}

	curr := l.head
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = new_node
}

func (l *LinkedList) Print_All() {
	curr := l.head
	for curr != nil {
		fmt.Print(curr.data, " -> ")
		curr = curr.next
	}
	fmt.Println("nil")
}
