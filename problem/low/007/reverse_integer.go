package reverse_integer

import (
	"math"
)

func ReverseInteger(x int) int {
	// 越界范围
	var upOutOfRange = int((math.Pow(2, 32) / 2) - 1)
	var downOutOfRange = int(0 - (math.Pow(2, 32) / 2))
	// 先确定正负数
	var flag = true
	if x < 0{
		if x < downOutOfRange {
			return 0
		}
		flag = false
		x = -x
	}else if x > upOutOfRange{
		return 0
	}

	// 将数字作为单个元素存储入组 倒序
	var baseNum []int

	var modNum = 10

	for{
		mod := x % modNum
		baseNum = append(baseNum, mod)
		x = (x - mod) / modNum
		if x < 1{
			break
		}
	}

	var res = 0
	// 将元素中的数一次抽取出来相加
	length := len(baseNum)
	for i := 0; i < length; i++ {
		if baseNum[i] == 0 {
			continue
		}
		res += baseNum[i] * int(math.Pow(float64(modNum), float64(length - i - 1)))
	}

	if !flag {
		res = 0 - res
		if res < downOutOfRange {
			return 0
		}
	}else if res > upOutOfRange{
		return 0
	}
	return int(res)
}
