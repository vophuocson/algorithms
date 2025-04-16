package maxpriorityqueue

func HeapifyDown(q *MaxPriorityQueue, i int) {
	lenQueue := q.size

	left := 2*i + 1
	right := 2*i + 2
	largest := i

	if left < lenQueue && q.data[largest].Priority < q.data[left].Priority {
		largest = left
	}

	if right < lenQueue && q.data[largest].Priority < q.data[right].Priority {
		largest = right
	}

	if i != largest {
		q.data[i], q.data[largest] = q.data[largest], q.data[i]
		HeapifyDown(q, largest)
	}
}

func HeapifyUp(q *MaxPriorityQueue, i int) {
	if i == 0 {
		return
	}
	parent := (i - 1) / 2
	if q.data[parent].Priority < q.data[i].Priority {
		q.data[parent], q.data[i] = q.data[i], q.data[parent]
		HeapifyUp(q, parent)
	}

}

func BuildMaxHeap(q *MaxPriorityQueue) {
	lenQ := q.Len()
	for i := lenQ - 1; i >= 0; i-- {
		HeapifyDown(q, i)
	}
}
