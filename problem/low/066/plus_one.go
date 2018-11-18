package plus_one

func plusOne(digits []int) []int {
	temp := []int{1}
	for i := len(digits) - 1; i >= 0; i--  {
		if digits[i] != 9{
			digits[i]++
			return digits
		}else {
			digits[i] = 0
			if i == 0 {
				digits = append(temp, digits...)
				return digits
			}
		}
	}
	return digits
}
