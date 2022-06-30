package linkedList

import "fmt"

// linked list : 1--->2--->3--->4--->5--->nil
//this is a singly linked List so the list can be traversed only one way

//define the node
// we use interface{} so i can use any type of data
type Node struct {
	data interface{}
	next *Node
}

//define the list
// use a pointer to the first node of the list

type List struct {
	size int
	head *Node
}

//function insert() that inserts a new node to the list
// the new node will be inserted in the last position

func (l *List) Insert(d interface{}) {
	// create the new node with the given data , and a pointer to the next node
	newNode := &Node{data: d, next: nil}
	if l.head == nil {
		l.head = newNode
	} else {
		current := l.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
	l.size += 1
}

//function size() that returns size of list

func SizeOf(l *List) int {
	return l.size
}

//function show() prints out the list

func (l *List) Show() {
	if l.head == nil {
		fmt.Println("Empty list")
	}
	current := l.head
	for current != nil {
		fmt.Printf("%v--->", current.data)
		current = current.next
	}
	fmt.Printf("nil\n")
}
