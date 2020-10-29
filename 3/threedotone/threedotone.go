package threedotone

// // type baseBuffer struct {
// // 	buffer [100]int
// // 	totalLen int
// // }
// const bufferMax = 100

// type Stack struct {
// 	baseBuffer       *[bufferMax]int
// 	inOutP           int
// 	len              int
// 	isLenOneElemNull bool
// 	index int
// }

// // var bb = new(baseBuffer)
// var bb = new([bufferMax]int)
// var stackPMap = make([]int, 3)
// var totalLen int
// var totalStackCount int

// func New() *Stack {
// 	inOutP := determineInOutP()
// 	s := &Stack{bb, inOutP, 0, true, totalStackCount}
// 	s.len = 1
// 	totalLen += 1
// 	stackPMap[s.index] = inOutP
// 	totalStackCount++
// 	return s
// }

// func determineInOutP() int {
// 	startP := stackPMap[0]
// 	currentInOutP := startP + totalLen
// 	if currentInOutP > bufferMax-1 {
// 		currentInOutP = totalLen - (totalLen - startP)
// 	}
// 	return currentInOutP
// }

// func (s *Stack) Push(value int) {
// 	if totalLen > bufferMax {
// 		return
// 	}
// 	//
// 	//s.baseBuffer[s.inOutP] = value
// 	if s.len == 1 && s.isLenOneElemNull {
// 		return
// 	}
// 	s.len++
// 	totalLen++
// }

// func (s *Stack) Pop() int {
// 	if s.len == 0 {
// 		return -1
// 	}
// 	popV := s.baseBuffer[s.inOutP]
// 	s.inOutP = nextPosition(s.inOutP)
// 	stackPMap[s.index] = s.inOutP
// 	s.len--
// 	totalLen--
// 	return popV
// }

// func nextPosition(currentP int) int {
// 	nextP := currentP+1
// 	if nextP >= bufferMax {
// 		nextP = 0
// 	}
// 	return nextP
// }

// func kthStackToEndDistance(index int) int {
// 	switch index {
// 	case 0:
// 		return totalLen
// 	case 1:
// 		return stackPMap[1] - 
// 	}
// }
