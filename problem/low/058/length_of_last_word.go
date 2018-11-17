package length_of_last_word

import (
	"strings"
)

func lengthOfLastWord(s string) int {
	s = strings.Trim(s, " ")
	if s == ""{
		return 0
	}

	ss := strings.Split(s, " ")

	return len(ss[len(ss) - 1])
}

func lengthOfLastWord2(s string) int {
	s = strings.Trim(s, " ")
	index := 0
	flag := true
	for i := len(s) - 1; i > 0; i-- {
		if s[i] == ' '{
			index = i + 1
			flag = false
			break
		}
	}
	if flag {
		return len(s)
	}
	return len(s) - index
}
