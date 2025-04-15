package heapsort

func Heapify(arr []int, size, i int) {
	left := 2*i + 1
	right := 2*i + 2
	largest := i
	if left < size && arr[left] > arr[largest] {
		largest = left
	}
	if right < size && arr[right] > arr[largest] {
		largest = right
	}
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		Heapify(arr, size, largest)
	}
}

func BuildHeapMax(arr []int) {
	lenArr := len(arr)
	for i := lenArr/2 - 1; i >= 0; i-- {
		Heapify(arr, lenArr, i)
	}
}

func HeapSort(arr []int) {
	lenArr := len(arr)
	BuildHeapMax(arr)
	for i := lenArr - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0)
	}
}
