package collections

type SSet[T any] interface {
	Size() int
	Add(x T) bool
	Remove(x T) bool
	Find(x T) (T, error)
}
