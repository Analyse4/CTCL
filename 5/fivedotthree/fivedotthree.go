package fivedotthree

import (
	"container/heap"

	mewheap "github.com/Analyse4/mew/adt/heap"
	mewbinary "github.com/Analyse4/mew/algs/binary"
)

// O(N0*N1^2)
func FlipBitToWin(num int) int {
	dc := mewbinary.GetCountsOfDigits(num)
	l := make([]int, 0)
	for i := 1; i <= dc; i++ {
		if !mewbinary.GetBit(num, i) {
			l = append(l, i)
		}
	}
	maxCount := dc - len(l) + 1
	for {
		for _, pos := range l {
			tmp := num
			mewbinary.SetBit(tmp, pos)
			posL := pos + 1
			posR := pos - 1
			countL := 0
			countR := 0
			for posL <= dc && mewbinary.GetBit(tmp, posL) {
				countL++
				posL++
			}
			for posR >= 1 && mewbinary.GetBit(tmp, posR) {
				countR++
				posR--
			}
			if countL+countR+1 == maxCount {
				return maxCount
			}
		}
		maxCount--
	}
}

// O(N-digits ~ N^2-digits)
func FlipBitToWinV2(num int) int {
	dc := mewbinary.GetCountsOfDigits(num)
	l := make([]int, 0)
	mh := mewheap.NewMax()
	for i := 1; i <= dc; i++ {
		if !mewbinary.GetBit(num, i) {
			l = append(l, i)
		}
	}
	for _, pos := range l {
		tmp := num
		mewbinary.SetBit(tmp, pos)
		posL := pos + 1
		posR := pos - 1
		countL := 0
		countR := 0
		for posL <= dc && mewbinary.GetBit(tmp, posL) {
			countL++
			posL++
		}
		for posR >= 1 && mewbinary.GetBit(tmp, posR) {
			countR++
			posR--
		}
		mh.Push(countL + countR + 1)
	}
	return heap.Pop(mh).(int)
}
