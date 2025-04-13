package maxsublist

import "math"

type subArrays struct {
	sum     int
	indices [][2]int
}

func FindMaxCrossingSubArrays(arr []int, left, mid, right int) subArrays {
	maxLeft := math.MinInt
	var idxLeft []int
	sum := 0

	for i := mid; i >= left; i-- {
		sum += arr[i]
		if sum > maxLeft {
			maxLeft = sum
			idxLeft = []int{i}
		} else if sum == maxLeft {
			idxLeft = append(idxLeft, i)
		}
	}
	maxRight := math.MinInt
	sum = 0
	var idxRight []int
	for i := mid + 1; i <= right; i++ {
		sum += arr[i]
		if sum > maxRight {
			maxRight = sum
			idxRight = []int{i}
		} else if sum == maxRight {
			idxRight = append(idxRight, i)
		}
	}
	var result subArrays
	result.sum = maxLeft + maxRight
	for _, li := range idxLeft {
		for _, ri := range idxRight {
			result.indices = append(result.indices, [2]int{li, ri})
		}
	}
	return result
}
