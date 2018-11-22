package same_binary

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func makeBinaryTree(index int, value []int) *TreeNode {
	length := len(value)
	t := &TreeNode{}
	t.Val = value[index - 1]

	if index + 1 < length && 2 * index <= length && 2 * index + 1<= length{
		t.Left = makeBinaryTree(2 * index, value)
		t.Right = makeBinaryTree(2 * index + 1, value)
	}
	return t
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil{
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left)&&isSameTree(p.Left, q.Left)
}
