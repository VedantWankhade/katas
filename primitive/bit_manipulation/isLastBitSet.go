package bitmanipulation

func IsLastBitSet(num int) bool {
	return (num & 1) == 1
}
