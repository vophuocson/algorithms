package maxpriorityqueue

import (
	"testing"
)

func TestHeapifyDown_MultipleCases(t *testing.T) {
	tests := []struct {
		name     string
		input    []*Item
		size     int
		startIdx int
		expected []int
	}{
		{
			name: "Already Max Heap",
			input: []*Item{
				{Value: "A", Priority: 10},
				{Value: "B", Priority: 5},
				{Value: "C", Priority: 3},
			},
			size:     3,
			startIdx: 0,
			expected: []int{10, 5, 3},
		},
		{
			name: "Swap once with left",
			input: []*Item{
				{Value: "A", Priority: 1},
				{Value: "B", Priority: 5},
				{Value: "C", Priority: 3},
			},
			size:     3,
			startIdx: 0,
			expected: []int{5, 1, 3},
		},
		{
			name: "Swap with right",
			input: []*Item{
				{Value: "A", Priority: 1},
				{Value: "B", Priority: 2},
				{Value: "C", Priority: 6},
			},
			size:     3,
			startIdx: 0,
			expected: []int{6, 2, 1},
		},
		{
			name: "Recursive swap down",
			input: []*Item{
				{Value: "A", Priority: 1},
				{Value: "B", Priority: 10},
				{Value: "C", Priority: 9},
				{Value: "D", Priority: 12},
				{Value: "E", Priority: 8},
			},
			size:     5,
			startIdx: 0,
			expected: []int{10, 12, 9, 1, 8},
		},
		{
			name: "Single element",
			input: []*Item{
				{Value: "A", Priority: 42},
			},
			size:     1,
			startIdx: 0,
			expected: []int{42},
		},
		{
			name:     "Empty heap",
			input:    []*Item{},
			size:     0,
			startIdx: 0,
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &MaxPriorityQueue{
				data: tt.input,
				size: tt.size,
			}

			HeapifyDown(q, tt.startIdx)

			for i, v := range tt.expected {
				if i >= len(q.data) || q.data[i].Priority != v {
					t.Errorf("Test %s: at index %d, expected Priority = %d, but got %d", tt.name, i, v, q.data[i].Priority)
				}
			}
		})
	}
}

func TestHeapifyUp(t *testing.T) {
	tests := []struct {
		name     string
		input    []*Item
		size     int
		index    int
		expected []int
	}{
		{
			name: "One swap with parent",
			input: []*Item{
				{Value: "A", Priority: 10},
				{Value: "B", Priority: 20}, // Index 1 -> parent is 0
			},
			size:  2,
			index: 1,
			expected: []int{
				20, 10,
			},
		},
		{
			name: "Multiple swaps to root",
			input: []*Item{
				{Value: "A", Priority: 10},
				{Value: "B", Priority: 5},
				{Value: "C", Priority: 15},
				{Value: "D", Priority: 2},
				{Value: "E", Priority: 30}, // Index 4 -> up to root
			},
			size:  5,
			index: 4,
			expected: []int{
				30, 10, 15, 2, 5,
			},
		},
		{
			name: "No swap needed (already valid max-heap)",
			input: []*Item{
				{Value: "A", Priority: 50},
				{Value: "B", Priority: 30},
				{Value: "C", Priority: 40},
			},
			size:  3,
			index: 2,
			expected: []int{
				50, 30, 40,
			},
		},
		{
			name: "Single element",
			input: []*Item{
				{Value: "A", Priority: 100},
			},
			size:     1,
			index:    0,
			expected: []int{100},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &MaxPriorityQueue{
				data: tt.input,
				size: tt.size,
			}

			HeapifyUp(q, tt.index)

			for i, v := range tt.expected {
				if i >= len(q.data) || q.data[i].Priority != v {
					t.Errorf("Test %s: at index %d, expected Priority = %d, but got %d", tt.name, i, v, q.data[i].Priority)
				}
			}
		})
	}
}

func TestExtract(t *testing.T) {
	tests := []struct {
		name     string
		input    []*Item
		expected []*Item // expected result after extraction
		size     int     // expected size of the heap after extraction
		max      *Item   // expected max item returned by Extract
	}{
		{
			name: "Multiple elements",
			input: []*Item{
				{Value: "A", Priority: 100},
				{Value: "B", Priority: 50},
				{Value: "C", Priority: 20},
			},
			expected: []*Item{
				{Value: "B", Priority: 50},
				{Value: "C", Priority: 20},
			},
			size: 2,
			max:  &Item{Value: "A", Priority: 100},
		},
		// Case 2: Heap chỉ có một phần tử
		{
			name: "Single element",
			input: []*Item{
				{Value: "A", Priority: 100},
			},
			expected: []*Item{},
			size:     0,
			max:      &Item{Value: "A", Priority: 100},
		},
		// Case 3: Heap rỗng ngay từ đầu
		{
			name:     "Empty heap",
			input:    []*Item{},
			expected: []*Item{},
			size:     0,
			max:      nil,
		},
		// Case 4: Sau khi gọi Extract 1 lần, heap chỉ còn 2 phần tử
		{
			name: "Extract until size 2",
			input: []*Item{
				{Value: "A", Priority: 100},
				{Value: "B", Priority: 50},
				{Value: "C", Priority: 20},
			},
			expected: []*Item{
				{Value: "B", Priority: 50},
				{Value: "C", Priority: 20},
			},
			size: 2,
			max:  &Item{Value: "A", Priority: 100},
		},
		// Case 5: Heap sau khi gọi Extract chỉ còn 1 phần tử
		{
			name: "Extract until size 1",
			input: []*Item{
				{Value: "A", Priority: 100},
			},
			expected: []*Item{},
			size:     0,
			max:      &Item{Value: "A", Priority: 100},
		},
		// Case 6: Heap có cùng một giá trị Priority, kiểm tra xem phần tử đầu tiên vẫn được trả về
		{
			name: "All same priority",
			input: []*Item{
				{Value: "A", Priority: 50},
				{Value: "C", Priority: 50},
				{Value: "B", Priority: 50},
			},
			expected: []*Item{
				{Value: "B", Priority: 50},
				{Value: "C", Priority: 50},
			},
			size: 2,
			max:  &Item{Value: "A", Priority: 50},
		},
		// Case 7: Heap với các phần tử giảm dần
		{
			name: "Descending order",
			input: []*Item{
				{Value: "A", Priority: 100},
				{Value: "B", Priority: 90},
				{Value: "C", Priority: 80},
			},
			expected: []*Item{
				{Value: "B", Priority: 90},
				{Value: "C", Priority: 80},
			},
			size: 2,
			max:  &Item{Value: "A", Priority: 100},
		},
		// Case 8: Heap với các phần tử tăng dần
		{
			name: "Ascending order",
			input: []*Item{
				{Value: "C", Priority: 30},
				{Value: "A", Priority: 10},
				{Value: "B", Priority: 20},
			},
			expected: []*Item{
				{Value: "B", Priority: 20},
				{Value: "A", Priority: 10},
			},
			size: 2,
			max:  &Item{Value: "C", Priority: 30},
		},
		// Case 9: Heap sau khi gọi Extract chỉ còn 1 phần tử, kiểm tra phần tử trả về đúng
		{
			name: "Extract until empty",
			input: []*Item{
				{Value: "A", Priority: 10},
			},
			expected: []*Item{},
			size:     0,
			max:      &Item{Value: "A", Priority: 10},
		},
	}

	// Lặp qua các test case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Khởi tạo MaxPriorityQueue từ input
			q := &MaxPriorityQueue{
				data: tt.input,
				size: len(tt.input),
			}

			// Gọi Extract và kiểm tra kết quả trả về
			max := q.Extract()

			// Kiểm tra max item trả về
			if max != nil && (max.Value != tt.max.Value || max.Priority != tt.max.Priority) {
				t.Errorf("Test %s: expected max %v, but got %v", tt.name, tt.max, max)
			}

			// Kiểm tra kích thước của heap
			if q.size != tt.size {
				t.Errorf("Test %s: expected heap size %d, but got %d", tt.name, tt.size, q.size)
			}

			// Kiểm tra các phần tử trong heap sau khi Extract
			for i, expectedItem := range tt.expected {
				if i >= len(q.data) || q.data[i].Priority != expectedItem.Priority || q.data[i].Value != expectedItem.Value {
					t.Errorf("Test %s: at index %d, expected item %v, but got %v", tt.name, i, expectedItem, q.data[i])
				}
			}
		})
	}
}
