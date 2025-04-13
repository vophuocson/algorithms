package maxsublist

import (
	"reflect"
	"testing"
)

func TestFindMaxCrossingSubArrays(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		left     int
		mid      int
		right    int
		expected subArrays
	}{
		{
			name:  "basic test",
			arr:   []int{2, -1, 3, -4, 5, 6, -2},
			left:  0,
			mid:   3,
			right: 6,
			expected: subArrays{
				sum:     11,
				indices: [][2]int{{0, 5}},
			},
		},
		{
			name:  "all negative",
			arr:   []int{-2, -3, -1, -4},
			left:  0,
			mid:   1,
			right: 3,
			expected: subArrays{
				sum:     -4,
				indices: [][2]int{{1, 2}},
			},
		},
		{
			name:  "single element crossing",
			arr:   []int{-2, 4, -1},
			left:  0,
			mid:   1,
			right: 2,
			expected: subArrays{
				sum:     3,
				indices: [][2]int{{1, 2}},
			},
		},
		{
			name:  "multiple crossing subarrays",
			arr:   []int{1, 2, -1, 3, 2},
			left:  0,
			mid:   2,
			right: 4,
			expected: subArrays{
				sum:     7,
				indices: [][2]int{{0, 4}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindMaxCrossingSubArrays(tt.arr, tt.left, tt.mid, tt.right)
			if got.sum != tt.expected.sum {
				t.Errorf("sum: expected %v, got %v", tt.expected.sum, got.sum)
			}
			if !reflect.DeepEqual(got.indices, tt.expected.indices) {
				t.Errorf("indices: expected %v, got %v", tt.expected.indices, got.indices)
			}
		})
	}
}
