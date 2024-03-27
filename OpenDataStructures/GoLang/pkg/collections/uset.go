package collections

type USet[T any] interface {
	Size() int
	Add(x T)
	Remove(x T) T
	Find(x T) T
}
