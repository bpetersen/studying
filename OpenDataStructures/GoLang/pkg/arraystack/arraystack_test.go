package arraystack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPushPopSize(t *testing.T) {
	stack := NewArrayStack[string]()

	stack.Push("A")
	stack.Push("B")
	stack.Push("C")
	stack.Push("D")
	stack.Push("E")

	var actual string

	assert.Equal(t, 5, stack.Size())

	actual, _ = stack.Pop()
	assert.Equal(t, "E", actual)
	assert.Equal(t, 4, stack.Size())

	actual, _ = stack.Pop()
	assert.Equal(t, "D", actual)
	assert.Equal(t, 3, stack.Size())

	actual, _ = stack.Pop()
	assert.Equal(t, "C", actual)
	assert.Equal(t, 2, stack.Size())

	actual, _ = stack.Pop()
	assert.Equal(t, "B", actual)
	assert.Equal(t, 1, stack.Size())

	actual, _ = stack.Pop()
	assert.Equal(t, "A", actual)
	assert.Equal(t, 0, stack.Size())

	_, err := stack.Pop()
	assert.ErrorContains(t, err, "Stack is empty")
}
