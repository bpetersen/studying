package binarytree

type Node[T any] struct {
	parent     *Node[T]
	data       T
	leftChild  *Node[T]
	rightChild *Node[T]
}

type TraversalType int

const (
	InOrder    TraversalType = iota //left, root, right
	PostOrder                       //left, right root
	PreOrder                        //root, left, right
	LevelOrder                      //BFS
)

type BinaryTree[T any] struct {
	root *Node[T]
}

func NewBinaryTree[T any]() *BinaryTree[T] {
	return &BinaryTree[T]{
		root: nil,
	}
}

type ForEach[T any] func(x T)

func (bt *BinaryTree[T]) ForEach(forEach ForEach[T], ordering TraversalType) {
	var previous *Node[T]
	var current *Node[T]
	var next *Node[T]
	current = bt.root
	for current != nil {
		if previous == current.parent {

			if ordering == PreOrder {
				forEach(current.data)
			}
			if current.leftChild != nil {
				next = current.leftChild
			} else if current.rightChild != nil {
				if ordering == InOrder {
					forEach(current.data)
				}
				next = current.rightChild
			} else {
				if ordering == InOrder || ordering == PostOrder {
					forEach(current.data)
				}
				next = current.parent
			}
		} else if previous == current.leftChild {
			if current.rightChild != nil {
				if ordering == InOrder {
					forEach(current.data)
				}
				next = current.rightChild
			} else {
				if ordering == PostOrder {
					forEach(current.data)
				}
				next = current.parent
			}
		} else {
			if ordering == PostOrder {
				forEach(current.data)
			}
			next = current.parent
		}
		previous = current
		current = next
	}
}

func forEachLevelOrder[T any](forEach ForEach[T], node *Node[T]) {
	//We're going to use a queue here
}
