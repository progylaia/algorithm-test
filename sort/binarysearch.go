package sort

import (
	"fmt"
)

func binarySearch(arr []int, data int) int {
	l, r := 0, len(arr)-1
	for l <= r {
		mid := (l + r) / 2
		if arr[mid] == data {
			return mid
		}
		if data > arr[mid] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}

func main()  {
	testArr := []int{2,3,5,7,9,12}
	index := binarySearch(testArr, 5)
	fmt.Printf("%d\n", index)
}