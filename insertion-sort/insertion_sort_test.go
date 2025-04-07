package insertionsort_test

import (
	insertionsort "algorithms/insertion-sort"
	"testing"
)

func Test_InsertionSort(t *testing.T) {
	for _, c := range []struct {
		arr    []int
		expect []int
	}{
		{
			arr:    []int{4, 2},
			expect: []int{2, 4},
		},
		{
			arr:    []int{1, 3, 2},
			expect: []int{1, 2, 3},
		},
		{
			arr:    []int{1},
			expect: []int{1},
		},
		{
			arr:    []int{5, 4, 3, 2, 1},
			expect: []int{1, 2, 3, 4, 5},
		},
	} {
		insertionsort.InsertionSort(c.arr)
		if !equal(c.arr, c.expect) {
			t.Errorf("InsertionSort(%v) = %v, want %v", c.arr, c.arr, c.expect)
		}
	}
}

func equal(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i := range arr1 {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}
