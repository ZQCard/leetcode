package implement_strstr

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"

)


type para struct {
	haystack string
	needle string
	output int
}

var params = []para{
	{//
		haystack:"hello",
		needle:"ll",
		output:2,
	},
	{
		haystack:"aaaaa",
		needle:"bba",
		output:-1,
	},
	{
		haystack:"card",
		needle:"",
		output:0,
	},
	{
		haystack:"mississippi",
		needle:"issip",
		output:4,
	},
}

func TestImplementStrstr(t *testing.T)  {
	for _, v := range params{
		Convey("实现strstr()", t, func() {
			So(implementStrstr(v.haystack, v.needle), ShouldEqual, v.output)
		})
	}
}