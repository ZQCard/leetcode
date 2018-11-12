package valid_parentheses

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

type para struct {
	input string
	output bool
}

var params = []para{
	{
		input:"()",
		output:true,
	},
	{
		input:"()[]{}",
		output:true,
	},
	{
		input:"(]",
		output:false,
	},
	{
		input:"([)]",
		output:false,
	},
	{
		input:"{[]}",
		output:true,
	},
}

func TestValidParentheses(t *testing.T){
	for _, v := range params{
		Convey("找出最大公共前缀", t, func() {
			So(validParentheses(v.input), ShouldEqual, v.output)
		})
	}
}
