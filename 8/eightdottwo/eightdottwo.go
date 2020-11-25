package eightdottwo

import "github.com/Analyse4/mew/adt/stack"

type position struct {
	x int
	y int
}

// func RobotSearchPath(r, c int, banPoint []*position, path []*position) (bool, []*position) {
// 	if r == 0 && c == 0 {
// 		path = append(path, &position{0, 0})
// 		return true, path
// 	}
// 	if isBanPoint(r, c, banPoint) {
// 		return false, path
// 	}
// 	isLead := false
// 	if c-1 >= 0 {
// 		isLead, path = RobotSearchPath(r, c-1, banPoint, path)
// 		if isLead {
// 			path = append(path, &position{r, c - 1})
// 		}
// 	}
// 	if c-1 < 0 || !isLead {
// 		isLead, path = RobotSearchPath(r-1, c, banPoint, path)
// 		if !isLead {
// 			return false, path
// 		}
// 		path = append(path, &position{r - 1, c})
// 	}
// 	return true, path
// }

// O(A(C+R)N), A is constant
func RobotSearchPath(r, c int, banPoint []*position, path *stack.Stack) (bool, *stack.Stack) {
	if r == 0 && c == 0 {
		return true, path
	}
	if isBanPoint(r, c, banPoint) {
		return false, path
	}
	isLead := false
	if c-1 >= 0 {
		isLead, path = RobotSearchPath(r, c-1, banPoint, path)
		if isLead {
			path.Push(&position{r, c - 1})
		}
	}
	if c-1 < 0 || !isLead {
		isLead, path = RobotSearchPath(r-1, c, banPoint, path)
		if !isLead {
			return false, path
		}
		path.Push(&position{r - 1, c})
	}
	return true, path
}

func isBanPoint(r, c int, banPoint []*position) bool {
	for _, p := range banPoint {
		if r == p.x && c == p.y {
			return true
		}
	}
	return false
}
