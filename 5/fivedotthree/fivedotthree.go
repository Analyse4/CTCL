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
	if num == 0 {
		return 1
	}
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

// Time: O(b)
// Space: O(b)
func FlipBitToWinV3(num int) int {
	cbl := countBitList(num)
	max := 0
	for i := 0; i < len(cbl); i += 2 {
		countL := 0
		countR := 0
		switch cbl[i] {
		case 0:
			countL = cbl[i+1]
		case 1:
			if i == len(cbl) {
				countL = 0
			} else {
				countL = cbl[i+1]
			}
			if i == 0 {
				countR = 0
			} else {
				countR = cbl[i-1]
			}
		default:
			if i == len(cbl) {
				countR = cbl[i-1]
			} else if i == 0 {
				countL = cbl[i+1]
			} else {
				if cbl[i+1] > cbl[i-1] {
					countL = cbl[i+1]
				} else {
					countR = cbl[i-1]
				}
			}
		}
		if max < countL+countR+1 {
			max = countL + countR + 1
		}
	}
	return max
}

// O(b)
func countBitList(num int) []int {
	l := make([]int, 0)
	targetBit := 0
	dc := mewbinary.GetCountsOfDigits(num)
	count := 0
	for i := 0; i < dc; i++ {
		if num&1 != targetBit {
			l = append(l, count)
			targetBit = 1 - targetBit
			count = 0
		}
		count++
		num >>= 1
	}
	l = append(l, count)
	return l
}
