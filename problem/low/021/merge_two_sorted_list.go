package merge_two_sorted_list

type ListNode struct {
	Value int
	Next  *ListNode
}

func makeListNode(nums []int) *ListNode {
	if len(nums) == 0{
		return &ListNode{}
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
	l1 := makeListNode([]int{1, 2, 3})
	l2 := makeListNode([]int{1, 3, 4})
	return mergeTwoSortedList(l1, l2)
}

func mergeTwoSortedList(l1 *ListNode, l2 *ListNode) *ListNode {
	// 申明一个结果链表
	res := make([]int, 0)

	for {
		// 比较l1和l2当前数的大小,小的放入res，并且指针后移
		if l1.Value <= l2.Value {
			res = append(res, l1.Value)
			l1 = l1.Next
		}

		if l2.Value <= l1.Value {
			res = append(res, l1.Value)
			l2 = l2.Next
		}

		if l1.Next == nil && l2.Next == nil {
			break
		}
	}
	return l1
}
