package mergesort

func Merge(arr []int, left, mid, right int) {
	nLeft := mid - left + 1
	nRight := right - mid

	arrLeft := make([]int, nLeft+1)
	arrRight := make([]int, nRight+1)

	for i := 0; i < nLeft; i++ {
		arrLeft[i] = arr[left+i]
	}
	for i := 0; i < nRight; i++ {
		arrRight[i] = arr[mid+1+i]
	}
	const INF = int(^uint(0) >> 1) // Max int
	arrLeft[nLeft] = INF
	arrRight[nRight] = INF

	var j, i = 0, 0
	for k := left; k <= right; k++ {
		if arrLeft[i] <= arrRight[j] {
			arr[k] = arrLeft[i]
			i++
		} else {
			arr[k] = arrRight[j]
			j++
		}
	}
}

func MergeSort(arr []int, left, right int) {
	if left < right {
		mid := (left + right) / 2
		MergeSort(arr, left, mid)
		MergeSort(arr, mid+1, right)
		Merge(arr, left, mid, right)
	}
}
