package fifo

// Element represents an element in the queue
type Element struct {
	Name  string
	Value int
}

// Queue represents all elements in the queue
type Queue struct {
	Elements *[]Element
}

// Enqueue puts an element in the queue
func (q *Queue) Enqueue(el *Element) (err error) {
	return nil
}

// Dequeue removes first element of the queue
func (q *Queue) Dequeue() (el *Element, err error) {
	return &Element{}, nil
}
