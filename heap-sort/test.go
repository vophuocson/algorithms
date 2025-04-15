package heapsort

import "fmt"

// Heapify function to maintain the heap property
func heapify(arr []int, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	// Check if left child is larger than root
	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	// Check if right child is larger than root
	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	// If the largest is not root, swap and continue heapifying
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, largest)
	}
}

// Run the heapify process to build the max-heap
func run() {
	arr := []int{9, 0, 3, 16, 13, 10, 1, 5, 7, 12, 4, 28, 27, 17}
	idxs := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	n := len(arr)

	// Build max heap
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
		fmt.Println(i)
		fmt.Printf("%v\n", arr)  // In ra mảng và chỉ số sau mỗi lần lặp
		fmt.Printf("%v\n", idxs) // In ra mảng và chỉ số sau mỗi lần lặp
	}

	// Print the resulting array after building the max-heap
	fmt.Println("Final Max-heap array:", arr)
}

func main() {
	run()
}
