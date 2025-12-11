package bitmanipulation

/*
Question:

	Return the parity of num (i.e (number of set bits) % 2)

Approach

	Count set bits of num, and return 1 if its odd, 0 if even

Complexity:

	Worst, Average Time: O(n); where n is number of bits in num
	Best Time: O(1)
*/
func ParityCountSetBits(num int) int {
	count := CountActiveBits(uint(num))
	return int(count % 2)
}

/*
Question:

	Return the parity of num (i.e (number of set bits) % 2)

Approach

	Count the number of times the lowest set bits in num can be dropped till the num becomes 0.

	Better average and best time complexity than counting set bits with brute force.

Complexity:

	Worst Time: O(n); where n is number of ALL bits in num
	Average Time: O(m); where m is number of SET bits in num. m < n
	Best Time: O(1)
*/
func ParityDropLowestSetBit(num int) int {
	parity := 0
	for num > 0 {
		// switch parity from 1 to 0 or 0 to 1 every time lowest set bit is dropped.
		// We can also increase a counter then return counter % 2 in the end instead.
		// but switching parity every time has the same effect
		num = DropLowestSetBit(num)
		parity ^= 1 // <- switch parity
	}
	return parity
}
