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

func (stack *Stack) Peek() (value interface{}) {
	value = stack.list.DeleteFirst()
	stack.list.InsertFirst(value)
	return
}

func (stack *Stack) IsEmpty() bool {
	return stack.list.IsEmpty()
}

func (stack *Stack) Display() {
	stack.list.Display()
}
