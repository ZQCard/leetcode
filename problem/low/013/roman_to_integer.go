package roman_to_integer

/*
	Symbol       Value
	I             1
	V             5
	X             10
	L             50
	C             100
	D             500
	M             1000
 */
 // 效率与最佳方法一致
func romanToInteger(s string) int {
	// 创建map表明罗马数与整数的关系
	m := make(map[string]int)
	m["I"] = 1
	m["V"] = 5
	m["X"] = 10
	m["L"] = 50
	m["C"] = 100
	m["D"] = 500
	m["M"] = 1000

	var res int
	for i :=  0; i < len(s); i++ {
		num, ok := m[string(s[i])]
		// 非法字符串
		if !ok {
			return 0
		}
		// 每个字母的数字相加
		// 处理特殊情况
		// IV XC CD
		if i + 1 == len(s){
			res += num
			break
		}
		if string(s[i]) == "I" && (string(s[i + 1]) == "V" || string(s[i + 1]) == "X"){
			i++
			num = m[string(s[i])] - 1
		}else if string(s[i]) == "X" && (string(s[i + 1]) == "L" || string(s[i + 1]) == "C"){
			i++
			num = m[string(s[i])] - 10
		}else if string(s[i]) == "C" && (string(s[i + 1]) == "D" || string(s[i + 1]) == "M"){
			i++
			num = m[string(s[i])] - 100
		}
		res += num
	}
	return res
}

func TheBest(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		cur := 0
		switch s[i] {
		case 'I':
			cur = 1
			if i < len(s)-1 {
				if s[i+1] == 'V' {
					cur = 4
					i++
					break
				}
				if s[i+1] == 'X' {
					cur = 9
					i++
					break
				}
			}
		case 'V':
			cur = 5
		case 'X':
			cur = 10
			if i < len(s)-1 {
				if s[i+1] == 'L' {
					cur = 40
					i++
					break
				}
				if s[i+1] == 'C' {
					cur = 90
					i++
					break
				}
			}
		case 'L':
			cur = 50
		case 'C':
			cur = 100
			if i < len(s)-1 {
				if s[i+1] == 'D' {
					cur = 400
					i++
					break
				}
				if s[i+1] == 'M' {
					cur = 900
					i++
					break
				}
			}
		case 'D':
			cur = 500
		case 'M':
			cur = 1000
		}
		res += cur
	}
	return res
}
