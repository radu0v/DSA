package doubleLinkedList

import "fmt"
// this is a double linked list implementation
//the list can be traversed both ways

//create the node. A pointer to a previous node, a pointer to the next node and the data

type Node struct{
	prev *Node
	data interface{}
	next *Node
}

// create the list. 
type List struct{
	size int
	head *Node	//pointer to first node
	tail *Node //pointer to last node
}

//------functions---------------------------------

// insert data into list at the last place
func (l *List) Insert(d interface{}){
	node := &Node(prev:nil,data:d,next:nil)
	if l.head == nil {
		l.head == node
		l.tail == node
	}else{
		current = l.head
		for current.next != nil{
			current = current.next
		}
		node.prev = current
		current.next = node
		l.tail = node
	}
	l.size += 1
}

func Show(l *List){
	if l.head == nil{
		fmt.Println("Empty list")
	}else{
		current := l.head
		for current != nil{
			fmt.Printf("%v--->",current.data)
			current = current.next
		}
		fmt.Println("nil")
	}
}

func SizeOf(l *List) int{
	return l.size
}

func showReverse(l *List){
	if l.head == nil{
		fmt.Println("Empty list")
	}else{
		current := l.tail
		for current != nil{
			fmt.Printf("%v--->",current.data)
			current = current.prev
		}
		fmt.Println("nil")
	}
}






