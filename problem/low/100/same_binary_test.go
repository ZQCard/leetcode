package same_binary

import (
	"fmt"
	"testing"
)

func TestSameBinary(t *testing.T) {
	value := []int{1, 2, 3, 4}
	p := makeBinaryTree(1, value)

	value2 := []int{1, 2, 3, 4}
	q := makeBinaryTree(1, value2)

	fmt.Println(isSameTree(p, q))
}
