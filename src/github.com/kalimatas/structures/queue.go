package structures

type Queue struct {
	list List
}

func (queue *Queue) Insert(value interface{}) {
	queue.list.InsertLast(value)
}

func (queue *Queue) Remove() (value interface{}) {
	value = queue.list.DeleteFirst()
	return
}

func (queue *Queue) IsEmpty() bool {
	return queue.list.IsEmpty()
}

func (queue *Queue) Display() {
	queue.list.Display()
}
