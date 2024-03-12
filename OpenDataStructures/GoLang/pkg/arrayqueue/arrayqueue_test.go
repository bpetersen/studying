package arrayqueue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {
	queue := NewArrayQueue[string]()

	queue.Enqueue("A")
	queue.Enqueue("B")
	queue.Enqueue("C")
	queue.Enqueue("D")
	queue.Enqueue("E")

	var actual string

	assert.Equal(t, 5, queue.Size())

	actual, _ = queue.Dequeue()
	assert.Equal(t, "A", actual)
	assert.Equal(t, 4, queue.Size())

	actual, _ = queue.Dequeue()
	assert.Equal(t, "B", actual)
	assert.Equal(t, 3, queue.Size())

	actual, _ = queue.Dequeue()
	assert.Equal(t, "C", actual)
	assert.Equal(t, 2, queue.Size())

	actual, _ = queue.Dequeue()
	assert.Equal(t, "D", actual)
	assert.Equal(t, 1, queue.Size())

	actual, _ = queue.Dequeue()
	assert.Equal(t, "E", actual)
	assert.Equal(t, 0, queue.Size())

	_, err := queue.Dequeue()
	assert.ErrorContains(t, err, "ArrayQueue is empty")
}
