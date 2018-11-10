package roman_to_integer

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)


type para struct {
	input string
	output int
}

var params = []para{
	{
		input:"III",
		output:3,
	},
	{
		input:"IV",
		output:4,
	},
	{
		input:"IX",
		output:9,
	},
	{
		input:"LVIII",
		output:58,
	},
	{
		input:"MCMXCIV",
		output:1994,
	},
}

func TestRomanToInteger(t *testing.T){
	for _, v := range params{
		Convey("判断整数是否是回文数", t, func() {
			So(romanToInteger(v.input), ShouldEqual, v.output)
		})
	}
}
