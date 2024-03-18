package dllist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSLListOperations(t *testing.T) {
	var list = NewDLList[string]()

	// Test Size initialize
	assert.Equal(t, 0, list.Size(), "list should be empty initially")

	//Test Add
	assert.NoError(t, list.Add(0, "B"), "Add operation failed")
	assert.Equal(t, 1, list.Size(), "list should contain one item after add")
	//B

	// Test Get
	val, err := list.Get(0)
	assert.NoError(t, err, "Get operation failed")
	assert.Equal(t, "B", val, "incorrect value retrieved by Get")

	//Test Add to front
	assert.NoError(t, list.Add(0, "A"), "Add operation failed")
	assert.Equal(t, 2, list.Size(), "list should contain two items")
	//A,B

	// Test Get
	val, err = list.Get(0)
	assert.NoError(t, err, "Get operation failed")
	assert.Equal(t, "A", val, "incorrect value retrieved by Get")

	val, err = list.Get(1)
	assert.NoError(t, err, "Get operation failed")
	assert.Equal(t, "B", val, "incorrect value retrieved by Get")

	//Test Add to back
	assert.NoError(t, list.Add(2, "D"), "Add operation failed")
	assert.Equal(t, 3, list.Size(), "list should contain three items")
	//A,B,D

	//Test Get
	val, err = list.Get(2)
	assert.NoError(t, err, "Get operation failed")
	assert.Equal(t, "D", val, "incorrect value retrieved by Get")

	//Test Add to back
	assert.NoError(t, list.Add(2, "C"), "Add operation failed")
	assert.Equal(t, 4, list.Size(), "list should contain three items")
	//A,B,C,D

	//Test Get
	val, err = list.Get(2)
	assert.NoError(t, err, "Get operation failed")
	assert.Equal(t, "C", val, "incorrect value retrieved by Get")

	// Test Set
	oldVal, err := list.Set(0, "W")
	assert.NoError(t, err, "Set operation failed")
	assert.Equal(t, "A", oldVal, "incorrect old value returned by Set")
	newVal, _ := list.Get(0)
	assert.Equal(t, "W", newVal, "incorrect value after Set operation")
	//W,B,C,D

	// Test Set
	oldVal, err = list.Set(1, "X")
	assert.NoError(t, err, "Set operation failed")
	assert.Equal(t, "B", oldVal, "incorrect old value returned by Set")
	newVal, _ = list.Get(1)
	assert.Equal(t, "X", newVal, "incorrect value after Set operation")
	//W,X,C,D

	// Test Set
	oldVal, err = list.Set(2, "Y")
	assert.NoError(t, err, "Set operation failed")
	assert.Equal(t, "C", oldVal, "incorrect old value returned by Set")
	newVal, _ = list.Get(2)
	assert.Equal(t, "Y", newVal, "incorrect value after Set operation")
	//W,X,Y,D

	// Test Set
	oldVal, err = list.Set(3, "Z")
	assert.NoError(t, err, "Set operation failed")
	assert.Equal(t, "D", oldVal, "incorrect old value returned by Set")
	newVal, _ = list.Get(3)
	assert.Equal(t, "Z", newVal, "incorrect value after Set operation")
	//W,X,Y,Z

	// Test Remove Front Half
	removedVal, err := list.Remove(1)
	assert.NoError(t, err, "Remove operation failed")
	assert.Equal(t, "X", removedVal, "incorrect value removed")
	assert.Equal(t, 3, list.Size(), "incorrect value after Remove operation")
	//W,Y,Z

	// Test Remove
	removedVal, err = list.Remove(1)
	assert.NoError(t, err, "Remove operation failed")
	assert.Equal(t, "Y", removedVal, "incorrect value removed")
	assert.Equal(t, 2, list.Size(), "incorrect value after Remove operation")
	//W,Z

	// Test Remove
	removedVal, err = list.Remove(1)
	assert.NoError(t, err, "Remove operation failed")
	assert.Equal(t, "Z", removedVal, "incorrect value removed")
	assert.Equal(t, 1, list.Size(), "incorrect value after Remove operation")
	//W

	// Test Remove
	removedVal, err = list.Remove(0)
	assert.NoError(t, err, "Remove operation failed")
	assert.Equal(t, "W", removedVal, "incorrect value removed")
	assert.Equal(t, 0, list.Size(), "incorrect value after Remove operation")

	//Test add and remove one more time
	assert.NoError(t, list.Add(0, "B"), "Add operation failed")
	assert.Equal(t, 1, list.Size(), "list should contain one item after add")

	// Test Remove
	removedVal, err = list.Remove(0)
	assert.NoError(t, err, "Remove operation failed")
	assert.Equal(t, "B", removedVal, "incorrect value removed")
	assert.Equal(t, 0, list.Size(), "incorrect value after Remove operation")

	_, err = list.Set(-1, "NOPE")
	assert.Error(t, err, "expected error for negative index in Set")

	err = list.Add(1, "NOPE")
	assert.Error(t, err, "expected error for out-of-bounds Add")

	_, err = list.Remove(0)
	assert.Error(t, err, "expected error for Remove on empty list")

	_, err = list.Get(-1)
	assert.Error(t, err, "expected error for negative index in Get")

	_, err = list.Get(0)
	assert.Error(t, err, "expected error for negative index in Get")
}

func TestGetNode(t *testing.T) {
	var dummyNode = node[string]{
		data: "DUMMY",
	}
	dummyNode.next = &dummyNode
	dummyNode.prev = &dummyNode
	var list = DLList[string]{
		dummy: &dummyNode,
		size:  3,
	}
	var cNode = node[string]{
		data: "C",
		next: &dummyNode,
		prev: &dummyNode,
	}
	dummyNode.next = &cNode
	dummyNode.prev = &cNode
	var bNode = node[string]{
		data: "B",
		next: &cNode,
		prev: &dummyNode,
	}
	cNode.prev = &bNode
	dummyNode.next = &bNode
	var aNode = node[string]{
		data: "A",
		next: &bNode,
		prev: &dummyNode,
	}

	bNode.prev = &aNode
	dummyNode.next = &aNode

	// next: dummy -> A -> B -> C -> dummy
	// prev: dummy -> C -> B -> A -> dummy

	actual := list.getNode(0)
	assert.Equal(t, "A", actual.data)

	actual = list.getNode(1)
	assert.Equal(t, "B", actual.data)

	actual = list.getNode(2)
	assert.Equal(t, "C", actual.data)
}

func TestGetAndAdd(t *testing.T) {
	var dummyNode = node[string]{
		data: "DUMMY",
	}
	dummyNode.next = &dummyNode
	dummyNode.prev = &dummyNode
	var list = DLList[string]{
		dummy: &dummyNode,
		size:  0,
	}
	list.Add(0, "C")
	valueAt, _ := list.Get(0)
	assert.Equal(t, "C", valueAt)

	list.Add(0, "B")
	valueAt, _ = list.Get(0)
	assert.Equal(t, "B", valueAt)

	list.Add(0, "A")
	valueAt, _ = list.Get(0)
	assert.Equal(t, "A", valueAt)

}
