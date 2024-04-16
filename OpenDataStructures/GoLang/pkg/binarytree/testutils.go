package binarytree

func buildBalancedBinaryTree() *BinaryTree[int] {
	//        3
	//       / \
	//      1   5
	//     / \ / \
	//    0  2 4  6

	rootNode := Node[int]{
		data: 3,
	}
	one := Node[int]{
		data:   1,
		parent: &rootNode,
	}
	five := Node[int]{
		data:   5,
		parent: &rootNode,
	}
	rootNode.leftChild = &one
	rootNode.rightChild = &five

	zero := Node[int]{
		data:   0,
		parent: &one,
	}
	two := Node[int]{
		data:   2,
		parent: &one,
	}
	one.leftChild = &zero
	one.rightChild = &two

	four := Node[int]{
		data:   4,
		parent: &five,
	}
	six := Node[int]{
		data:   6,
		parent: &five,
	}
	five.leftChild = &four
	five.rightChild = &six

	binaryTree := BinaryTree[int]{
		root: &rootNode,
	}

	return &binaryTree
}
