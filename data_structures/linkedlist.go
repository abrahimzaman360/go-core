package data_structures

import (
	"encoding/json"
)

type Node struct {
	data map[string]any
	next *Node
	prev *Node
}

type LinkedList struct {
	head *Node
	tail *Node
	size int
}

// {
// 	"item1": "string",
// 	"fog": {
// 		[
// 			"Hello"
// 		]
// 	}
// }

func (l *LinkedList) Append(val map[string]any) {
	new_node := &Node{data: val}

	if l.head == nil {
		l.head = new_node
		l.tail = new_node
		l.size += 1
		return
	}

	curr := l.head
	for curr.next != nil {
		curr = curr.next
	}

	curr.next = new_node
	new_node.prev = l.tail
	l.tail = new_node
	l.size += 1
}

func (l *LinkedList) Print_All() (any, error) {
	var result []map[string]interface{}
	curr := l.head
	for curr != nil {
		result = append(result, curr.data)
		curr = curr.next
	}

	jsonData, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

func (l *LinkedList) Return_Size() int {
	return l.size
}
