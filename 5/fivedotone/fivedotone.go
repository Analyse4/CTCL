package fivedotone

// O(bit count of m)
func Insertion(n, m, i, j int) int {
	mBitCount := 0
	endNum := 1
	for m >= endNum {
		mBitCount++
		endNum <<= 1
	}
	if m == 0 {
		mBitCount = 1
	}
	mask := -1<<j | (1<<(i-1) - 1)
	return n&mask | (m << (j - mBitCount + 1))
}
