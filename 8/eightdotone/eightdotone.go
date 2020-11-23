package eightdotone

// // Time: O(N)	Space: O(N)
// func ThreeStep(n int) int {
// 	if n == 0 {
// 		return 0
// 	}
// 	if n == 1 {
// 		return 1
// 	}
// 	if n == 2 {
// 		return 2
// 	}
// 	if n == 3 {
// 		return 4
// 	}
// 	return ThreeStep(n-3) * 4
// }

// // Time: O(N)	Space: O(1)
// func ThreeStepV2(n int) int {
// 	if n == 0 {
// 		return 0
// 	}
// 	if n == 1 {
// 		return 1
// 	}
// 	if n == 2 {
// 		return 2
// 	}
// 	if n == 3 {
// 		return 4
// 	}
// 	currentStep := 3
// 	currentMCou := 4
// 	for {
// 		if currentStep+3 <= n {
// 			currentStep += 3
// 			currentMCou *= 4
// 			if currentStep == n {
// 				return currentMCou
// 			}
// 		} else if currentStep+2 <= n {
// 			currentStep += 2
// 			currentMCou *= 2
// 			if currentStep == n {
// 				return currentMCou
// 			}
// 		} else {
// 			return currentMCou
// 		}
// 	}
// }

// O(3^n)
func ThreeStepV3(n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	return ThreeStepV3(n-1) + ThreeStepV3(n-2) + ThreeStepV3(n-3)
}

// Time: O(N)	Space: O(N)
func ThreeStepV4(n int, l []int) int {
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	if l[n] != 0 {
		return l[n]
	}
	l[n] = ThreeStepV4(n-1, l) + ThreeStepV4(n-2, l) + ThreeStepV4(n-3, l)
	return l[n]
}

func ThreeStepV5(n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	a := 1
	b := 0
	c := 0
	tc := 0
	for i := 1; i <= n; i++ {
		tc = a + b + c
		c = b
		b = a
		a = tc
	}
	return tc
}
