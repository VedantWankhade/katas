package bitmanipulation

/*
Question:

	Count set bits in positive num

Approach:

	Right shift each bit in num by 1; so that the 1 bit to left of lsb becomes the next lsb,
	then bitwise and the num with 1.
	If result is 1 then increment the coundt Repeat till num becomes 1.

Complexity:

	Worst, Average, Best Time: O(n); where n is number of bits in num
*/
func CountActiveBits(num uint) uint {
	var count uint = 0
	for num > 0 {
		count += num & 1
		num >>= 1
	}
	return count
}
