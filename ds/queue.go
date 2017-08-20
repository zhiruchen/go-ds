package ds

type Queue struct {
	Elements []interface{}
}

func (q *Queue) Enqueue(e interface{}) {
	q.Elements = append(q.Elements, e)
}

func (q *Queue) IsEmpty() bool {
	return len(q.Elements) > 0
}

func (q *Queue) Dequeue() interface{} {
	if len(q.Elements) == 0 {
		return nil
	}
	e := q.Elements[0]
	q.Elements = q.Elements[1:]
	return e
}
