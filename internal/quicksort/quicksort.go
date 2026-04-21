package quicksort

func partition(arr []int, low int, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j <= high-1; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func Sort(arr []int, low int, high int) {
	if low < high {
		pi := partition(arr, low, high)
		Sort(arr, low, pi-1)
		Sort(arr, pi+1, high)
	}
}
