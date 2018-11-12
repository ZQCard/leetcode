package merge_two_sorted_list

type ListNode struct {
	Value int
	Next  *ListNode
}

func makeListNode(nums []int) *ListNode {
	if len(nums) == 0{
		return nil
	}
	res := &ListNode{
		Value: nums[0],
	}
	temp := res
	for i := 1; i < len(nums); i++ {
		temp.Next = &ListNode{Value: nums[i]}
		temp = temp.Next
	}
	return res
}

func runMergeTwoSortedList() *ListNode {
	l1 := makeListNode([]int{1, 5, 6})
	l2 := makeListNode([]int{})
	return mergeTwoSortedList(l1, l2)
}

func mergeTwoSortedList(l1 *ListNode, l2 *ListNode) *ListNode {
	// 申明一个结果链表
	res := make([]int, 0)

	for {
		// 两者均不为nil
		if l1 != nil && l2 != nil{
			valueOne := l1.Value
			valueTwo := l2.Value
			if valueOne <= valueTwo {
				res = append(res, valueOne)
				l1 = l1.Next
			} else {
				res = append(res, valueTwo)
				l2 = l2.Next
			}
		}

		// l1为nil
		if l1 == nil && l2 != nil{
			res = append(res, l2.Value)
			l2 = l2.Next
		}

		// l2为nil
		if l2 == nil && l1 != nil{
			res = append(res, l1.Value)
			l1 = l1.Next
		}

		if l1 == nil && l2 == nil{
			 break
		}
	}
	return makeListNode(res)
}
