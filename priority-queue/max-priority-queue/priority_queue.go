package maxpriorityqueue

import (
	"errors"
	"math"
)

type Item struct {
	Value    string
	Priority int
}

type MaxPriorityQueue struct {
	data []*Item
	size int
}

func (q *MaxPriorityQueue) Len() int {
	return q.size
}

func (q *MaxPriorityQueue) Maximum() *Item {
	return q.data[0]
}

func (q *MaxPriorityQueue) Extract() *Item {
	if q.size <= 0 {
		return nil
	}
	q.size = q.size - 1
	max := q.data[0]
	if q.size > 0 {
		q.data[0] = q.data[q.size]
		HeapifyDown(q, 0)
	}
	return max
}

func (q *MaxPriorityQueue) IncreasePriority(i, priority int) error {
	if i >= q.size {
		return errors.New("invalid")
	}
	if priority <= q.data[i].Priority {
		return errors.New("invalid")
	}
	q.data[i].Priority = priority
	HeapifyUp(q, i)
	return nil
}

func (q *MaxPriorityQueue) Insert(i *Item) {
	q.size = q.size + 1
	q.data = append(q.data, &Item{Priority: math.MinInt, Value: i.Value})
	q.IncreasePriority(q.size-1, i.Priority)
}
