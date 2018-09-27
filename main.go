package main

import (
	"fmt"
	"github.com/algorithm-test/sort"
)

func main() {
	arr := []int{4,8,11,1,22,12,5,32,6,9,21,18}
	sort.SelectionSort(arr, len(arr))
	fmt.Println(arr)
}