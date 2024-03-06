package collections

type Queue[T any] interface {
	Enqueue(x T)
	Dequeue() (T, error)
	Size() int
}
