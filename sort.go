package dsa

func BubbleSort(list []int) []int {
	for i := 0; i < len(list)-1; i++ {
		for j := i + 1; j < len(list); j++ {
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
	return list
}

func SelectionSort(list []int) []int {
	for lo := 0; lo < len(list)-1; lo++ {
		max := lo + 1
		for i := lo + 2; i < len(list); i++ {
			if list[i] < list[max] {
				max = i
			}
		}
		list[lo], list[max] = list[max], list[lo]
	}
	return list
}

func InsertionSort(list []int) []int {
	for i := 1; i < len(list); i++ {
		key := list[i]
		j := i - 1
		for ; j >= 0 && key < list[j]; j-- {
			list[j+1] = list[j]
		}
		list[j+1] = key
	}
	return list
}

func HeapSort(list []int) []int {
	heap := NewHeap[int](func(i1, i2 int) int {
		return i1 - i2
	})
	for _, i := range list {
		heap.add(i)
	}
	for i := range list {
		list[i] = *heap.get()
	}
	return list
}

func merge(left, right []int) []int {
	i, j := 0, 0
	aux := make([]int, 0)

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			aux = append(aux, left[i])
			i++
		} else {
			aux = append(aux, right[j])
			j++
		}
	}
	if i >= len(left) {
		for ; j < len(right); j++ {
			aux = append(aux, right[j])
		}
	} else {
		for ; i < len(left); i++ {
			aux = append(aux, left[i])
		}
	}
	return aux
}

func MergeSort(list []int) []int {
	if len(list) < 2 {
		return list
	}
	left := list[0 : len(list)/2]
	right := list[len(list)/2:]
	left = MergeSort(left)
	right = MergeSort(right)
	return merge(left, right)
}

func partition(list []int, lo, hi int) ([]int, int) {
	pivotIndex := hi
	pivot := list[pivotIndex]
	i, j := lo, lo
	for j < hi {
		if list[j] < pivot {
			list[i], list[j] = list[j], list[i]
			i++
		}
		j++
	}
	list[i], list[pivotIndex] = pivot, list[i]
	return list, i
}

func quickSort(list []int, lo, hi int) []int {
	if hi-lo+1 < 2 {
		return list
	}
	list, pivotIndex := partition(list, 0, hi)
	list = quickSort(list, 0, pivotIndex-1)
	list = quickSort(list, pivotIndex+1, hi)
	return list
}

func QuickSort(list []int) []int {
	return quickSort(list, 0, len(list)-1)
}
