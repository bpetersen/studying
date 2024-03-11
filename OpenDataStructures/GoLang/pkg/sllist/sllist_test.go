package sllist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSLListOperations(t *testing.T) {
	var list = NewSLList[string]()

	// Test Size initially
	assert.Equal(t, 0, list.Size(), "list should be empty initially")

	//Test Add
	assert.NoError(t, list.Add(0, "B"), "Add operation failed")
	assert.Equal(t, 1, list.Size(), "list should contain one item after add")

	// Test Get
	val, err := list.Get(0)
	assert.NoError(t, err, "Get operation failed")
	assert.Equal(t, "B", val, "incorrect value retrieved by Get")

	//Test Add to front
	assert.NoError(t, list.Add(0, "A"), "Add operation failed")
	assert.Equal(t, 2, list.Size(), "list should contain two items")

	// Test Get
	val, err = list.Get(0)
	assert.NoError(t, err, "Get operation failed")
	assert.Equal(t, "A", val, "incorrect value retrieved by Get")

	//Test Add to back
	assert.NoError(t, list.Add(1, "C"), "Add operation failed")
	assert.Equal(t, 3, list.Size(), "list should contain three items")

	// Test Get
	val, err = list.Get(2)
	assert.NoError(t, err, "Get operation failed")
	assert.Equal(t, "C", val, "incorrect value retrieved by Get")

	// Test Set
	oldVal, err := list.Set(0, "X")
	assert.NoError(t, err, "Set operation failed")
	assert.Equal(t, "A", oldVal, "incorrect old value returned by Set")
	newVal, _ := list.Get(0)
	assert.Equal(t, "X", newVal, "incorrect value after Set operation")

	// Test Set
	oldVal, err = list.Set(1, "Y")
	assert.NoError(t, err, "Set operation failed")
	assert.Equal(t, "B", oldVal, "incorrect old value returned by Set")
	newVal, _ = list.Get(1)
	assert.Equal(t, "Y", newVal, "incorrect value after Set operation")

	// Test Set
	oldVal, err = list.Set(2, "Z")
	assert.NoError(t, err, "Set operation failed")
	assert.Equal(t, "C", oldVal, "incorrect old value returned by Set")
	newVal, _ = list.Get(2)
	assert.Equal(t, "Z", newVal, "incorrect value after Set operation")

	// Test Remove
	removedVal, err := list.Remove(1)
	assert.NoError(t, err, "Remove operation failed")
	assert.Equal(t, "Y", removedVal, "incorrect value removed")
	assert.Equal(t, 2, list.Size(), "list should be empty after remove")

	// Test Remove
	removedVal, err = list.Remove(1)
	assert.NoError(t, err, "Remove operation failed")
	assert.Equal(t, "Z", removedVal, "incorrect value removed")
	assert.Equal(t, 1, list.Size(), "list should be empty after remove")

	// Test Remove
	removedVal, err = list.Remove(0)
	assert.NoError(t, err, "Remove operation failed")
	assert.Equal(t, "X", removedVal, "incorrect value removed")
	assert.Equal(t, 0, list.Size(), "list should be empty after remove")

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
