package arraylist

const DEFAULT_INITIAL_CAPACITY = 10

type ArrayList[T any] struct {
	data []T
	size int
	head int
}

func NewArrayList[T any](initialCapacity ...int) *ArrayList[T] {
	capacity := DEFAULT_INITIAL_CAPACITY
	if len(initialCapacity) > 0 {
		capacity = initialCapacity[0]
	}

	return &ArrayList[T]{
		data: make([]T, capacity),
		size: 0,
		head: 0,
	}
}
