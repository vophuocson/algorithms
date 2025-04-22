package selection

import (
	"math/rand"
	"time"
)

func RandomIntInRange(min, max int) int {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	return r.Intn(max-min+1) + min
}

func Partition(arr []int, left, right int) int {
	idxPivot := RandomIntInRange(left, right)
	arr[idxPivot], arr[right] = arr[right], arr[idxPivot]
	pivot := arr[right]
	mid := left - 1
	for i := left; i <= right-1; i++ {
		if arr[i] <= pivot {
			mid++
			arr[mid], arr[i] = arr[i], arr[mid]
		}
	}
	arr[mid+1], arr[right] = arr[right], arr[mid+1]
	return mid + 1
}

func RandomizedSelect(arr []int, left, right, i int) int {
	if left == right {
		return arr[right]
	}
	partitionIndex := Partition(arr, left, right)
	k := partitionIndex - left + 1
	if k == i {
		return arr[partitionIndex]
	}
	if k > i {
		return RandomizedSelect(arr, left, partitionIndex-1, i)
	} else {
		return RandomizedSelect(arr, partitionIndex+1, right, i-k)
	}
}
