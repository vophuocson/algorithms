package maxsublist

func MaxSubArrays(arr []int, low, high int) subArrays {
	if high == low {
		return subArrays{sum: arr[low], indices: [][2]int{{low, high}}}
	} else {
		var result subArrays
		mid := (low + high) / 2
		maxLeft := MaxSubArrays(arr, low, mid)
		maxRight := MaxSubArrays(arr, mid+1, high)
		maxCrossing := FindMaxCrossingSubArrays(arr, low, mid, high)
		maxSum := Max(maxLeft.sum, maxRight.sum, maxCrossing.sum)
		result.sum = maxSum
		if maxLeft.sum == maxSum {
			result.indices = append(result.indices, maxLeft.indices...)
		}
		if maxRight.sum == maxSum {
			result.indices = append(result.indices, maxRight.indices...)
		}
		if maxCrossing.sum == maxSum {
			result.indices = append(result.indices, maxCrossing.indices...)
		}
		return result
	}
}

func Max(a, b, c int) int {
	if a >= b && a >= c {
		return a
	}
	if b >= a && b >= c {
		return b
	}
	return c
}
