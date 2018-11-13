package merge_two_sorted_list

import (
	"fmt"
	"testing"
)

func TestMergeTwoSortedList(t *testing.T)  {
	res := runMergeTwoSortedList()
	for {
		if res == nil{
			break
		}
		fmt.Printf("%v-->", res.Value)
		res = res.Next
	}
}
