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
