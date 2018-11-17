package maximum_subarray

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)


type para struct {
	haystack []int
	output int
}

var params = []para{
	{//
		haystack:[]int{-2, -1},
		output:-1,
	},
	{
		haystack:[]int{-2,1,-3,4,-1,2,1,-5,4},
		output:6,
	},
}

func TestMaximumSubArray(t *testing.T) {
	for _, v := range params{
		Convey("最大子数组和", t, func() {
			So(TheBest(v.haystack), ShouldEqual, v.output)
		})
	}
}
