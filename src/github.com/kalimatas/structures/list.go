package structures

import "fmt"

// Element of the list.
type node struct {
	value interface{}
	next  *node
}

func (n *node) String() string {
	return fmt.Sprintf("%s", n.value)
}

// List data structure.
type List struct {
	first *node
	last  *node
}

func (list *List) IsEmpty() bool {
	return list.first == nil
}

// InsertFirst inserts a node in the beginning of a list.
func (list *List) InsertFirst(value interface{}) {
	newNode := node{value: value}
	if list.IsEmpty() {
		list.last = &newNode
	}

	newNode.next = list.first
	list.first = &newNode
}

// InsertLast inserts a node to the end of a list.
func (list *List) InsertLast(value interface{}) {
	newNode := node{value: value}
	if list.IsEmpty() {
		list.first = &newNode
	} else {
		list.last.next = &newNode
	}
	
	list.last = &newNode
}

// DeleteFirst removes and returns the first node of a list.
func (list *List) DeleteFirst() (value interface{}) {
	if list.IsEmpty() {
		return
	}

	value = list.first.value
	if list.first.next == nil {
		list.last = nil
	}
	
	list.first = list.first.next
	return
}

// Display outputs the elements of a list.
func (list *List) Display() {
	for current := list.first; current != nil; current = current.next {
		fmt.Printf("%s ", current)
	}
}
