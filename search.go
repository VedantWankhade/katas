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
