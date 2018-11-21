package merge_sorted_array

import "testing"

func TestMergeSortArray(t *testing.T)  {
	num1 := []int{4,5,6,0,0,0}
	num2 := []int{2,5,6}
	TheBest(num1, 3, num2, 3)
}