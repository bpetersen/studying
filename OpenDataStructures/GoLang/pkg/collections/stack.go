package collections

type Stack[T any] interface {
	Push(x T)
	Pop() (T, error)
	Size() T
}
