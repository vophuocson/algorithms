package heapsort_test

import (
	heapsort "algorithms/heap-sort"
	"reflect"
	"testing"
)

func TestHeapify(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		i        int
		size     int
		expected []int
	}{
		{
			name:     "Heapify root index 0",
			input:    []int{3, 9, 2, 1, 4, 5},
			i:        0,
			size:     6,
			expected: []int{9, 4, 2, 1, 3, 5}, // max heap for root
		},
		{
			name:     "Heapify middle node index 1",
			input:    []int{10, 1, 2, 20, 30},
			i:        1,
			size:     5,
			expected: []int{10, 30, 2, 20, 1}, // 30 bubble up
		},
		{
			name:     "No change needed",
			input:    []int{50, 30, 40},
			i:        0,
			size:     3,
			expected: []int{50, 30, 40},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arr := make([]int, len(tt.input))
			copy(arr, tt.input)
			heapsort.Heapify(arr, tt.size, tt.i)
			if !reflect.DeepEqual(arr, tt.expected) {
				t.Errorf("got %v, want %v", arr, tt.expected)
			}
		})
	}
}

func TestBuildHeapMax(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int // một cây max-heap hợp lệ
	}{
		{
			name:     "Basic heap",
			input:    []int{3, 9, 2, 1, 4, 5},
			expected: []int{9, 4, 5, 1, 3, 2},
		},
		{
			name:     "Already max heap",
			input:    []int{100, 90, 80, 70, 60},
			expected: []int{100, 90, 80, 70, 60},
		},
		{
			name:     "Reverse order",
			input:    []int{1, 2, 3, 4, 5, 6, 7},
			expected: []int{7, 5, 6, 4, 2, 1, 3},
		},
		{
			name:     "Empty array",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Single element",
			input:    []int{42},
			expected: []int{42},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arr := make([]int, len(tt.input))
			copy(arr, tt.input)
			heapsort.BuildHeapMax(arr)
			if !isValidMaxHeap(arr) {
				t.Errorf("BuildHeapMax failed. got %v, which is not a valid max heap", arr)
			}
			// Optional: test against expected if needed
			if len(tt.expected) > 0 && !reflect.DeepEqual(arr, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, arr)
			}
		})
	}
}

// Helper: kiểm tra một mảng có là max-heap hay không
func isValidMaxHeap(arr []int) bool {
	n := len(arr)
	for i := 0; i < n; i++ {
		left := 2*i + 1
		right := 2*i + 2
		if left < n && arr[i] < arr[left] {
			return false
		}
		if right < n && arr[i] < arr[right] {
			return false
		}
	}
	return true
}

func TestHeapSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Normal array",
			input:    []int{4, 10, 3, 5, 1},
			expected: []int{1, 3, 4, 5, 10},
		},
		{
			name:     "Already sorted",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Reverse order",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "With duplicates",
			input:    []int{2, 3, 2, 1, 4},
			expected: []int{1, 2, 2, 3, 4},
		},
		{
			name:     "Single element",
			input:    []int{42},
			expected: []int{42},
		},
		{
			name:     "Empty array",
			input:    []int{},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arr := make([]int, len(tt.input))
			copy(arr, tt.input)
			heapsort.HeapSort(arr)
			if !reflect.DeepEqual(arr, tt.expected) {
				t.Errorf("HeapSort() = %v, want %v", arr, tt.expected)
			}
		})
	}
}
