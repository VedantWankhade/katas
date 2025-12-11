package bitmanipulation

/*
Question:

	Check if the lsb in num is 1

Approach:

	Bitwise and of num with 1 will return 1 if lsb in num is 1

Complexity:

	Worst, Average, Best Time: O(1)
*/
func IsLastBitSet(num int) bool {
	return (num & 1) == 1
}
