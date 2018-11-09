package reverse_integer

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

type para struct {
	input int
	output int32
}

var params = []para{
	{
		input:123,
		output:321,
	},
	{
		input:-123,
		output:-321,
	},
	{
		input:1563847412,
		output:0,
	},
	{
		input:1534236469,
		output:0,
	},
	{
		input:-2147483648,
		output:0,
	},
}

func TestReverseInteger(t *testing.T) {
	for _, v := range params{
		Convey("测试整数翻转", t, func() {
			So(ReverseInteger(v.input), ShouldEqual, v.output)
		})
	}
}
