package binarytree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinaryTreeConstructor(t *testing.T) {
	binaryTree := NewBinaryTree[string]()

	assert.NotNil(t, binaryTree)
}

func TestBinaryTreeInOrderTraversal(t *testing.T) {

	actual := []int{}
	forEach := func(i int) {
		actual = append(actual, i)
	}

	binaryTree := buildBalancedBinaryTree()

	binaryTree.ForEach(forEach, InOrder)
	assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6}, actual)
}

func TestBinaryTreePostOrderTraversal(t *testing.T) {

	actual := []int{}
	forEach := func(i int) {
		actual = append(actual, i)
	}

	binaryTree := buildBalancedBinaryTree()

	binaryTree.ForEach(forEach, PostOrder)
	assert.Equal(t, []int{0, 2, 1, 4, 6, 5, 3}, actual)
}

func TestBinaryTreePreOrderTraversal(t *testing.T) {

	actual := []int{}
	forEach := func(i int) {
		actual = append(actual, i)
	}

	binaryTree := buildBalancedBinaryTree()

	binaryTree.ForEach(forEach, PreOrder)
	assert.Equal(t, []int{3, 1, 0, 2, 5, 4, 6}, actual)
}
