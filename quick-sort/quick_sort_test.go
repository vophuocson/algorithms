package quicksort

import "testing"

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
