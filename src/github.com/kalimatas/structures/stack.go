package structures

type Stack struct {
	list List
}

func (stack *Stack) Push(value interface{}) {
	stack.list.InsertFirst(value)
}

func (stack *Stack) Pop() (value interface{}) {
	value = stack.list.DeleteFirst()
	return
}

func (stack *Stack) Display() {
	stack.list.Display()
}
