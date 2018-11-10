package palindrome_number

func palindromeNumber(x int) bool {
	// 排除小于零的情况
	if x < 0 {
		return false
	}
	// 存放整数位数
	var temp []int
	for {
		mod := x % 10
		temp = append(temp, mod)
		x = x / 10
		if x < 1 {
			break
		}
	}
	length := len(temp)
	middle := length / 2

	for i := 0; i < middle; i++ {
		if temp[i] == temp[length-1-i] {
			continue
		} else {
			return false
		}
	}
	return true
}
