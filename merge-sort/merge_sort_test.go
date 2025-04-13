package mergesort_test

import (
	mergesort "algorithms/merge-sort"
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		input    []int
		left     int
		mid      int
		right    int
		expected []int
	}{
		{
			input:    []int{1, 3, 5, 2, 4, 6},
			left:     0,
			mid:      2,
			right:    5,
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			input:    []int{2, 4, 6, 1, 3, 5},
			left:     0,
			mid:      2,
			right:    5,
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			input:    []int{1, 2, 3, 4, 5, 6},
			left:     0,
			mid:      2,
			right:    5,
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			input:    []int{5, 6, 7, 1, 2, 3},
			left:     0,
			mid:      2,
			right:    5,
			expected: []int{1, 2, 3, 5, 6, 7},
		},
	}

	for _, tt := range tests {
		arrCopy := make([]int, len(tt.input))
		copy(arrCopy, tt.input)
		mergesort.Merge(arrCopy, tt.left, tt.mid, tt.right)

		if !reflect.DeepEqual(arrCopy, tt.expected) {
			t.Errorf("Merge() failed. Got %v, want %v", arrCopy, tt.expected)
		}
	}
}

func TestMergeSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{5, 2, 4, 6, 1, 3},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			input:    []int{9, 7, 5, 3, 1},
			expected: []int{1, 3, 5, 7, 9},
		},
		{
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			input:    []int{3, 3, 2, 1, 2},
			expected: []int{1, 2, 2, 3, 3},
		},
		{
			input:    []int{},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		arrCopy := make([]int, len(tt.input))
		copy(arrCopy, tt.input)

		if len(arrCopy) > 0 {
			mergesort.MergeSort(arrCopy, 0, len(arrCopy)-1)
		}
		if !reflect.DeepEqual(arrCopy, tt.expected) {
			t.Errorf("MergeSort(%v) = %v, expected %v", tt.input, arrCopy, tt.expected)
		}
	}
}
