package _4_queue

type Node struct {
	Info string
	Next *Node
}

type Queue struct {
	First *Node
	Last  *Node
}

func (q *Queue) IsEmpty() bool {
	return q.First == nil
}

func (q *Queue) Enqueue(s string) {
	t := &Node{Info: s, Next: nil}
	if q.IsEmpty() {
		// Empty Queue
		q.First = t
	} else {
		q.Last.Next = t
		q.Last = t
	}
}

func (q *Queue) Dequeue() (string, bool) {
	if q.IsEmpty() {
		q.Last = nil
		return "", false
	} else {
		r := q.First.Info
		q.First = q.First.Next
		return r, true
	}
}
