package eightdottwelve

import (
	"math"
)

// type Position struct {
// 	x int
// 	y int
// }

// type limit struct {
// 	row      map[int]bool
// 	cow      map[int]bool
// 	diagonal []*Position
// }

// func EightQueen(n int, allForms [][]*Position) [][]*Position {
// 	if n == 1 {
// 		for i := 0; i < 64; i++ {
// 			tmp := make([]*Position, 0)
// 			allForms = append(allForms, tmp)
// 		}
// 		index := 0
// 		for r := 0; r < 8; r++ {
// 			for c := 0; c < 8; c++ {
// 				allForms[index] = append(allForms[index], &Position{r, c})
// 				index++
// 			}
// 		}
// 		return allForms
// 	}
// 	allForms = EightQueen(n-1, allForms)
// 	extendAllForms := make([][]*Position, 0)
// 	allFormLimit := make([]*limit, 0)

// 	for _, form := range allForms {
// 		row := make(map[int]bool, 8)
// 		cow := make(map[int]bool, 8)
// 		diagonal := make([]*Position, 0)
// 		for _, p := range form {
// 			row[p.x] = true
// 			cow[p.y] = true
// 			tmpUpR := p.x
// 			tmpUpC := p.y
// 			for tmpUpR != 0 && tmpUpC != 0 {
// 				tmpUpR += 1
// 				tmpUpC -= 1
// 				diagonal = append(diagonal, &Position{tmpUpR, tmpUpC})
// 			}
// 			tmpDownR := p.x
// 			tmpDownC := p.y
// 			for tmpDownR != 7 && tmpDownC != 7 {
// 				tmpDownR -= 1
// 				tmpDownC += 1
// 				diagonal = append(diagonal, &Position{tmpDownR, tmpDownC})
// 			}
// 			tmpUpRR := p.x
// 			tmpUpRC := p.y
// 			for tmpUpRR != 0 && tmpUpRC != 0 {
// 				tmpUpRR += 1
// 				tmpUpRC += 1
// 				diagonal = append(diagonal, &Position{tmpUpRR, tmpUpRC})
// 			}
// 			tmpDownRR := p.x
// 			tmpDownRC := p.y
// 			for tmpDownRR != 7 && tmpDownRC != 7 {
// 				tmpDownRR -= 1
// 				tmpDownRC -= 1
// 				diagonal = append(diagonal, &Position{tmpDownRR, tmpDownRC})
// 			}
// 		}
// 		lim := &limit{row, cow, diagonal}
// 		allFormLimit = append(allFormLimit, lim)
// 	}

// 	for i, lim := range allFormLimit {
// 		for r := 0; r < 8; r++ {
// 			if lim.row[r] {
// 				continue
// 			}
// 			for c := 0; c < 8; c++ {
// 				if lim.cow[c] {
// 					continue
// 				}
// 				np := &Position{r, c}
// 				flag := true
// 				for _, tmpp := range lim.diagonal {
// 					if tmpp.x == np.x && tmpp.y == np.y {
// 						flag = false
// 						break
// 					}
// 				}
// 				if flag {
// 					extendAllForms = append(extendAllForms, append(allForms[i], np))
// 				}
// 			}
// 		}
// 	}
// 	return extendAllForms
// }

const maxQueen = 8

func EightQueenV2(n int, poss []int, allPoss [][]int) [][]int {
	if n == maxQueen {
		allPoss = append(allPoss, poss)
	} else {
		for i := 0; i < 8; i++ {
			if checkValid(n, i, poss) {
				poss[n] = i
				allPoss = EightQueenV2(n+1, poss, allPoss)
			}
		}
	}
	return allPoss
}

func checkValid(row, cow int, poss []int) bool {
	for i := 0; i < row; i++ {
		if cow == poss[i] {
			return false
		}
		if float64(row-i) == math.Abs(float64(cow)-float64(poss[i])) {
			return false
		}
	}
	return true
}

// func Tslice(l []int) {
// 	l = append(l, 4)
// }
