package fivedottwo

import (
	mewbinary "github.com/Analyse4/mew/algs/binary"
)

// O(number of decimal places)
func BinaryToString(bin float64) string {
	s := "0."
	for bin != float64(int64(bin)) {
		if len(s) >= 32 {
			return "ERROR"
		}
		bin *= 2
		if mewbinary.GetBit(int(bin), 1) {
			s += "1"
		} else {
			s += "0"
		}
	}
	return s
}
