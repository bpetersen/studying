package collections

type Deque[T any] interface {
	Size() int

	AddFirst(x T)
	AddLast(x T)

	RemoveFirst() (T, error)
	RemoveLast() (T, error)
}
