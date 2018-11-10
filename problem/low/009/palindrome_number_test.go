package palindrome_number

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)


type para struct {
	input int
	output bool
}

var params = []para{
	{
		input:123,
		output:false,
	},
	{
		input:-123,
		output:false,
	},
	{
		input:123321,
		output:true,
	},
	{
		input:12321,
		output:true,
	},
}

func TestPalindromeNumber(t *testing.T)  {
	for _, v := range params{
		Convey("判断整数是否是回文数", t, func() {
			So(palindromeNumber(v.input), ShouldEqual, v.output)
		})
	}
}
