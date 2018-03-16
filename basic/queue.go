package basic

type Queue struct {
	elements *LinkedList
	count int32
}

func (q *Queue) Enqueue(data interface{}) {

}

func (q *Queue) Dequeue() (interface{}, error) {
	return nil, nil
}

func (q *Queue) IsEmpty() bool {
	return q.count == 0
}