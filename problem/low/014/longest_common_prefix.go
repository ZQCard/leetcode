package longest_common_prefix

// 与最优解效率相同
func longestCommonPrefix(strs []string) string {
	// 排除空数组
	if len(strs) == 0{
		return ""
	}
	// 找出最短长度的元素 以此为基准
	var length = len(strs[0])
	for i := 0; i < len(strs); i++ {
		if len(strs[i]) < length {
			length = len(strs[i])
		}
	}
	// 存放相同字符
	var str = ""

	// 循环传入的数组
	var flag = true
	for i := 0; i < length; i++ {
		var temp = strs[0][i]
		for j := 1; j < len(strs); j++ {
			if temp != strs[j][i] {
				flag = false
				break
			}
		}
		if !flag {
			break
		}
		str += string(temp)
	}
	return str
}