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
}

// InsertFirst inserts a node in the beginning of a list.
func (list *List) InsertFirst(value interface{}) {
	newNode := node{value, list.first}
	list.first = &newNode
}

// DeleteFirst removes and returns the first node of a list.
func (list *List) DeleteFirst() (value interface{}) {
	value = list.first.value
	list.first = list.first.next
	return
}

// Display outputs the elements of a list.
func (list *List) Display() {
	for current := list.first; current != nil; current = current.next {
		fmt.Printf("%s ", current)
	} 
}
