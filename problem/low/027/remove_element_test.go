package remove_element

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)


type para struct {
	input []int
	val int
	output int
}

var params = []para{
	{
		input:[]int{3,2,2,3},
		val:2,
		output:2,
	},
	{
		input:[]int{0,1,2,2,3,0,4,2},
		val:2,
		output:5,
	},
}

func TestRemoveElement(t *testing.T)  {
	for _, v := range params{
		Convey("找出最大公共前缀", t, func() {
			So(removeElement(v.input, v.val), ShouldEqual, v.output)
		})
	}
}