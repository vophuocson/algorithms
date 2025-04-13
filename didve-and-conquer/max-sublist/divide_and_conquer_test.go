package maxsublist

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMaxSubArrays(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		expected subArrays
	}{
		{
			name: "simple case",
			arr:  []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			expected: subArrays{
				sum:     6,
				indices: [][2]int{{3, 6}}, // 4 + -1 + 2 + 1 = 6
			},
		},
		{
			name: "all negative",
			arr:  []int{-4, -3, -2, -1},
			expected: subArrays{
				sum:     -1,
				indices: [][2]int{{3, 3}},
			},
		},
		{
			name: "all positive",
			arr:  []int{1, 2, 3, 4},
			expected: subArrays{
				sum:     10,
				indices: [][2]int{{0, 3}},
			},
		},
		{
			name: "multiple max subarrays",
			arr:  []int{1, -1, 1, -1, 1},
			expected: subArrays{
				sum:     1,
				indices: [][2]int{{0, 0}, {2, 2}, {0, 2}, {4, 4}, {2, 4}, {0, 4}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MaxSubArrays(tt.arr, 0, len(tt.arr)-1)
			fmt.Println(result.indices)
			if result.sum != tt.expected.sum {
				t.Errorf("expected sum %d, got %d", tt.expected.sum, result.sum)
			}
			if !reflect.DeepEqual(result.indices, tt.expected.indices) {
				t.Errorf("expected indices %v, got %v", tt.expected.indices, result.indices)
			}
		})
	}
}
