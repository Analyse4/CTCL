package tendotone

// TIME: O(N)
// SPACE: O(1)
func Merge(a, b []int, length int) []int {
	aIndex := length - 1
	bIndex := len(b) - 1
	tarIndex := len(a) - 1
	for {
		if a[aIndex] > b[bIndex] {
			a[tarIndex] = a[aIndex]
			aIndex--
		} else {
			a[tarIndex] = b[bIndex]
			bIndex--
		}
		tarIndex--
		if aIndex < 0 {
			for i := bIndex; i >= 0; i-- {
				a[tarIndex] = b[i]
				tarIndex--
			}
			return a
		}
		if bIndex < 0 {
			return a
		}
	}
}
