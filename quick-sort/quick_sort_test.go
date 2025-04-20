package quicksort

import (
	"testing"
)

func TestPartition(t *testing.T) {
	tests := []struct {
		arr         []int
		left, right int
		expected    []int
	}{
		{
			arr:      []int{3, 2, 5, 1, 6},
			left:     0,
			right:    4,
			expected: []int{3, 2, 5, 1, 6},
		},
		{
			arr:      []int{5, 1, 4, 2, 3},
			left:     0,
			right:    4,
			expected: []int{1, 2, 3, 5, 4},
		},
		{
			arr:      []int{7, 6, 5, 8, 9},
			left:     0,
			right:    4,
			expected: []int{7, 6, 5, 8, 9},
		},
		{
			arr:      []int{1, 2, 3, 4, 5},
			left:     0,
			right:    4,
			expected: []int{1, 2, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		// Copy the array to prevent modifying the original array during tests
		arr := make([]int, len(tt.arr))
		copy(arr, tt.arr)

		// Call the Partition function
		Partition(arr, tt.left, tt.right)

		// Check if the result is correct
		for i := range arr {
			if arr[i] != tt.expected[i] {
				t.Errorf("For arr %v, expected %v, but got %v", tt.arr, tt.expected, arr)
				break
			}
		}
	}
}

func TestQuickSort(t *testing.T) {
	tests := []struct {
		arr      []int
		expected []int
	}{
		{
			arr:      []int{3, 2, 5, 1, 6},
			expected: []int{1, 2, 3, 5, 6},
		},
		{
			arr:      []int{5, 1, 4, 2, 3},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			arr:      []int{7, 6, 5, 8, 9},
			expected: []int{5, 6, 7, 8, 9},
		},
		{
			arr:      []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			arr:      []int{5, 5, 5, 5, 5},
			expected: []int{5, 5, 5, 5, 5},
		},
	}

	for _, tt := range tests {
		arr := make([]int, len(tt.arr))
		copy(arr, tt.arr)

		QuickSort(arr, 0, len(arr)-1)

		for i := range arr {
			if arr[i] != tt.expected[i] {
				t.Errorf("For arr %v, expected %v, but got %v", tt.arr, tt.expected, arr)
				break
			}
		}
	}
}

func TestQuicksortWithHoarePartition(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Unsorted array",
			input:    []int{5, 3, 8, 4, 2, 7, 1, 10},
			expected: []int{1, 2, 3, 4, 5, 7, 8, 10},
		},
		{
			name:     "Already sorted",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Reverse sorted",
			input:    []int{9, 8, 7, 6, 5},
			expected: []int{5, 6, 7, 8, 9},
		},
		{
			name:     "All elements equal",
			input:    []int{7, 7, 7, 7},
			expected: []int{7, 7, 7, 7},
		},
		{
			name:     "With duplicates",
			input:    []int{4, 2, 4, 3, 2},
			expected: []int{2, 2, 3, 4, 4},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			arr := make([]int, len(tc.input))
			copy(arr, tc.input)

			QuicksortWithHoarePartition(arr, 0, len(arr)-1)

			if !equal(arr, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, arr)
			}
		})
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
