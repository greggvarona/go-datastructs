package sorting

func QuickSort(a []int) {
	quickSort(a, 0, len(a)-1)
}

func quickSort(a []int, startIndex int, endIndex int) {
	if endIndex-startIndex <= 0 {
		return
	}
	mid := partition(a, startIndex, endIndex)
	quickSort(a, startIndex, mid-1) // quick sort left
	quickSort(a, mid+1, endIndex)   // quick sort right
}

func partition(a []int, startIndex int, endIndex int) int {
	pivot := a[endIndex]

	for endIndex > startIndex {
		if a[endIndex-1] > pivot {
			a[endIndex] = a[endIndex-1]
			endIndex--
		} else {
			temp := a[endIndex-1]
			a[endIndex-1] = a[startIndex]
			a[startIndex] = temp
			startIndex++
		}
	}
	a[endIndex] = pivot
	return endIndex
}
