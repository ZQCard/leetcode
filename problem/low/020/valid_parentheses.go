package valid_parentheses

import (
	"strings"
)

// 与最优解区别在于 死板的运用字符串 而不知道用byte 导致运行速度变慢
// 而核心思路是一样的
func validParentheses(s string) bool {
	m := make(map[string]string)
	m["["] = "]"
	m["("] = ")"
	m["{"] = "}"
	m["]"] = "["
	m[")"] = "("
	m["}"] = "{"

	temp := ""

	for i := 0; i < len(s); i++ {
		// 有非法字符
		if strings.Index("[]{}()", string(s[i])) == -1 {
			return false
		}
		// 如过字串为0直接加上
		// 如果当前字符匹配temp最后一个字符
		// 成功 删除temp最后一个字符
		// 失败 temp加上当前字符
		if len(temp) == 0 || (m[string(s[i])] != string(temp[len(temp)-1])) {
			temp += string(s[i])
		} else {
			temp = temp[0: len(temp) - 1]
		}
	}
	if len(temp) > 0{
		return false
	}
	return true
}

func TheBest(s string) bool {
	var m = map[byte]byte{')': '(', '}': '{', ']': '['}
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		elem, ok := m[s[i]]
		if ok {
			if len(stack) > 0 && stack[len(stack)-1] == elem {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
