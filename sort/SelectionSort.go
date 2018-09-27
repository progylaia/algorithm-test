package sort

func SelectionSort(arr []int, length int) {
	for i := 0; i < length; i++ {
		minIndex := i
		for j := i + 1; j < length;j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}
