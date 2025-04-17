package quicksort

func Partition(arr []int, left, right int) int {
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
