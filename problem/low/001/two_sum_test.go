package two_sum

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

type param struct {
	num []int
	target int
	expect int
}

var para = []param{
	{
		num:[]int{3, 2, 4},
		target:6,
		expect:2,
	},
	{
		num:[]int{3, 2, 4},
		target:8,
		expect:0,
	},
	{
		num:nil,
		target:8,
		expect:0,
	},
}

func TestTwoSum(t *testing.T) {
	for _, v := range para{
		Convey("自己实现:", t, func() {
			So(twoSumTwoMap(v.num, v.target), ShouldHaveLength, v.expect)
		})
	}
}
