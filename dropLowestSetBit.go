package katas

/*
Question:

	Drop the lowest set bit in num, and return the resulting integer

Approach:

	Bitwise and num with num - 1.
	num - 1 carries a carry from the lowest set bit to all the way till next (rightmost) set bit, inverting the bits between.
	and then the next bitwise and with num reverts all the inverted bits, leaving us with all the same bits as in num except the original lowest set bit

Complexity:

	Worst, Average, Best Time: O(1)
*/
func DropLowestSetBit(num int) int {
	return num & (num - 1)
}
