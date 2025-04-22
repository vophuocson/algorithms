package quicksort

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
	x := arr[right]
	i := left - 1
	for j := left; j <= right-1; j++ {
		if arr[j] <= x {
			i = i + 1
			arr[j], arr[i] = arr[i], arr[j]
		}
	}
	arr[i+1], arr[right] = arr[right], arr[i+1]
	return i + 1
}

func QuickSort(arr []int, left, right int) {
	if left < right {
		partition := Partition(arr, left, right)
		QuickSort(arr, left, partition-1)
		QuickSort(arr, partition+1, right)
	}
}

func HoarePartition(arr []int, left, right int) int {
	indexRandom := RandomIntInRange(left, right)
	arr[indexRandom], arr[left] = arr[left], arr[indexRandom]
	pivot := arr[left]
	i := left - 1
	j := right + 1

	for {
		for {
			i++
			if arr[i] >= pivot {
				break
			}
		}
		for {
			j--
			if arr[j] <= pivot {
				break
			}
		}
		if i >= j {
			return j
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func QuicksortWithHoarePartition(arr []int, left, right int) {
	if left < right {
		partition := HoarePartition(arr, left, right)
		QuicksortWithHoarePartition(arr, left, partition)
		QuicksortWithHoarePartition(arr, partition+1, right)
	}
}
