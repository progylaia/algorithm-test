package sort

func QuickSort(arr []int, n int) {
	qiuckSort3Ways(arr, 0, n-1)
}

func quickSort(arr []int, l, r int) {
	if l >= r {
		return
	}
	p := partition2Ways(arr, l, r)
	quickSort(arr, l, p-1)
	quickSort(arr, p+1, r)
}

func qiuckSort3Ways(arr []int, l, r int)  {
	if l >= r {
		return
	}
	lt, gt := partition3Ways(arr, l, r)
	quickSort(arr, l, lt)
	quickSort(arr, gt, r)
}

//对arr[l...r]切分,返回p使的arr[l...p-1] < arr[p]，arr[p+1...r] > arr[p]
func partition(arr []int, l, r int) int {
	v := arr[l]
	//arr[l+1...j] < arr[v]，arr[j+1...i-1] > arr[v]
	j := l
	for i := l + 1; i <= r; i++ {
		if arr[i] < v {
			arr[i], arr[j+1] = arr[j+1], arr[i]
			j++
		}
	}
	arr[l], arr[j] = arr[j], arr[l]
	return j
}

func partition2Ways(arr []int, l, r int) int {
	v := arr[l]
	//arr[l+1...i) < arr[v]，arr(j...r] > arr[v]
	i := l + 1
	j := r
	for {
		for i <= r && arr[i] < v {
			i++
		}
		for j >= l+1 && arr[j] > v {
			j--
		}
		if i > j {
			break
		}
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	arr[l], arr[j] = arr[j], arr[l]
	return j
}

func partition3Ways(arr []int, l, r int) (lt, gt int) {
	v := arr[l]
	//arr[l+1...lt] < v, arr[lt+1...i) = v, arr[gt...r] > v
	lt = l
	gt = r + 1
	i := l + 1
	for i < gt {
		if arr[i] < v {
			arr[i], arr[lt+1] = arr[lt+1], arr[i]
			lt++
			i++
		} else if arr[i] > v {
			arr[i], arr[gt-1] = arr[gt-1], arr[i]
			gt--
		} else {
			i++
		}
	}
	arr[l], arr[lt] = arr[lt], arr[l]
	lt--
	return
}
