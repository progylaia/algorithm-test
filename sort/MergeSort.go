package sort

func MergeSort(arr []int, n int) {
	mergeSort(arr, 0, n-1)
}

//[l...r]区间进行归并排序
func mergeSort(arr []int, l, r int) {
	if l >= r {
		return
	}
	mid := (l + r) / 2
	mergeSort(arr, l, mid)
	mergeSort(arr, mid+1, r)
	if arr[mid] > arr[mid+1] {
		merge(arr, l, mid, r)
	}
}

//[l...mid],[mid+1, r]
func merge(arr []int, l, mid, r int) {
	var aux []int
	for i := l; i <= r; i++ {
		aux = append(aux, arr[i])
	}
	i := l
	j := mid + 1
	for k := l; k <= r;k++ {
		if i > mid {
			arr[k] = aux[j-l]
			j++
		} else if j > r {
			arr[k] = aux[i-l]
			i++
		} else if aux[i-l] < aux[j-l] {
			arr[k] = aux[i-l]
			i++
		} else if aux[j-l] < aux[i-l] {
			arr[k] = aux[j-l]
			j++
		}
	}
}
