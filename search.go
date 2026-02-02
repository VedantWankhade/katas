package dsa

func BinarySearch[T any](key T, list []T, comp func(T, T) int) int {
	lo, hi := 0, len(list)-1

	for lo <= hi {
		mid := (hi + lo) / 2
		if comp(key, list[mid]) == 0 {
			return mid
		}
		if comp(key, list[mid]) < 0 {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return -1
}

func binSearchRec[T any](key T, list []T, lo, hi int, comp func(T, T) int) int {
	if hi-lo+1 <= 0 {
		return -1
	}
	if hi-lo+1 == 1 && comp(key, list[lo]) == 0 {
		return lo
	}
	if hi-lo+1 == 1 && comp(key, list[lo]) != 0 {
		return -1
	}

	mid := (lo + hi) / 2
	if comp(key, list[mid]) == 0 {
		return mid
	}
	if comp(key, list[mid]) < 0 {
		return binSearchRec(key, list, lo, mid-1, comp)
	}
	return binSearchRec(key, list, mid+1, hi, comp)
}

func BinarySearchRec[T any](key T, list []T, comp func(T, T) int) int {
	return binSearchRec(key, list, 0, len(list)-1, comp)
}
