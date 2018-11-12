package longest_common_prefix

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

type para struct {
	input []string
	output string
}

var params = []para{
	{
		input:[]string{"flower","flow","flight"},
		output:"fl",
	},
	{
		input:[]string{"dog","race","car"},
		output:"",
	},
	{
		input:[]string{},
		output:"",
	},
}

func TestRomanToInteger(t *testing.T){
	for _, v := range params{
		Convey("找出最大公共前缀", t, func() {
			So(longestCommonPrefix(v.input), ShouldEqual, v.output)
		})
	}
}