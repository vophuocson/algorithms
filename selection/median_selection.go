package selection

func GetMedian(arr []int) int {
	l := len(arr)
	for i := 1; i < l; i++ {
		k := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > k {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = k
	}
	return arr[len(arr)/2]
}

func GetPivot(arr []int) int {
	var medians []int
	l := len(arr)
	if l <= 5 {
		return GetMedian(arr)
	}
	for i := 0; i < l; i += 5 {
		var median int
		if l-i >= 5 {
			median = GetMedian(arr[i : i+5])
		} else {
			median = GetMedian(arr[i:])
		}
		medians = append(medians, median)
	}
	return GetPivot(medians)
}
