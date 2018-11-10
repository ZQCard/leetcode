package reverse_integer

import (
	"math"
)

func ReverseInteger(x int) int {
	// 先确定正负数
	var flag = true
	if x < 0{
		flag = false
		x = -x
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
		if res < math.MinInt32 {
			return 0
		}
	}else if res > math.MaxInt32{
		return 0
	}
	return int(res)
}

func ReverseIntegerTwo(x int) int {
	sign := 1

	// 处理负数
	if x < 0 {
		sign = -1
		x = -1 * x
	}

	res := 0
	for x > 0 {
		// 取出x的末尾
		temp := x % 10
		// 放入 res 的开头
		res = res*10 + temp
		// x 去除末尾
		x = x / 10
	}

	// 还原 x 的符号到 res
	res = sign * res

	// 处理 res 的溢出问题
	if res > math.MaxInt32 || res < math.MinInt32 {
		res = 0
	}

	return res
}

func TheBest(x int) int{
	var digits []int8

	for i := x; i != 0; i = i / 10{
		digits = append(digits, int8(i % 10))
	}

	var res int
	for i := len(digits) - 1; i >= 0; i--{
		res += int(digits[i]) * int(math.Pow10(len(digits) - i - 1))
	}
	if res > math.MaxInt32 || res < math.MinInt32{
		return 0
	}
	return res
}
