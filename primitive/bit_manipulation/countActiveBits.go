package bitmanipulation

// count set bits in possitive num
//
// worst time O(n); n is number of bits in num
func CountActiveBits(num uint) uint {
	var count uint = 0
	for num > 0 {
		count += num & 1
		num >>= 1
	}
	return count
}
