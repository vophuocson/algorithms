package queue

type Node struct {
	Value int
	Next  *Node
}

type Queue struct {
	Head *Node
	Tail *Node
	L    int
}

func (q *Queue) Enqueue(node *Node) {
	if q.Tail != nil {
		q.Tail.Next = node
		q.Tail = q.Tail.Next
	} else {
		q.Tail = node
		q.Head = node
	}
	q.L += 1
}

func (q *Queue) Dequeue() *Node {
	if q.Head == nil {
		return nil
	}
	node := q.Head
	q.Head = q.Head.Next
	q.L -= 1
	if q.Head == nil {
		q.Tail = nil
	}
	return node
}
