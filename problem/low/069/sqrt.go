package sqrt

import "math"

func sqrt(x int) int {
	return int(math.Sqrt(float64(x)))
}

func TheBest(x int) int {
	if x == 0 || x == 1 {
		return x
	}

	left, right := 0, x

	for left < right {
		mid := (right+left)/2

		if left*left <= x && (left+1)*(left+1) > x {
			return left
		}

		if mid*mid <= x {
			left = mid
		} else {
			right = mid
		}
	}

	return 0
}