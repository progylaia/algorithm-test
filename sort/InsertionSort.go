package sort

func InsertionSort1(arr []int, n int) {
	for i := 1; i < n; i++ {
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
}

func InsertionSort2(arr []int, n int)  {
	for i := 1;i < n ;i++  {
		ele := arr[i]
		var j int
		for j = i;j > 0 && ele < arr[j-1]; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = ele
	}
}
