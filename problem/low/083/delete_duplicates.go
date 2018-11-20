package remove_deplicates_from_sorted_list

type ListNode struct {
	Val int
	Next *ListNode
}

func run(nums []int) *ListNode {
	list := makeList(nums)
	list = deleteDuplicates(list)
	return list
}

func makeList(nums []int) *ListNode {
	if len(nums) == 0 {
		return &ListNode{}
	}
	list := &ListNode{
		Val:nums[0],
	}
	temp := list
	for i := 1; i < len(nums); i++ {
		temp.Next = &ListNode{Val:nums[i]}
		temp = temp.Next
	}
	return list
}
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	nums := []int{}
	for {
		current := head.Val
		if head != nil{
			if head.Next != nil && current == head.Next.Val{
				head.Next = head.Next.Next
				continue
			}
		}
		head = head.Next
		nums = append(nums, current)
		if head == nil{
			break
		}
	}
	return makeList(nums)
}

func TheBest(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var node *ListNode = head
	for node != nil {
		if node.Next != nil && node.Val == node.Next.Val {
			node.Next = node.Next.Next
		} else {
			node = node.Next
		}
	}
	return head
}