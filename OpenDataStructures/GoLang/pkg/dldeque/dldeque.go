package dldeque

import (
	"fmt"
	"open_data_structures/pkg/collections"
	"open_data_structures/pkg/dllist"
)

type DLDeque[T any] struct {
	data *dllist.DLList[T]
}

var _ collections.Deque[int] = &DLDeque[int]{}

func NewDLDeque[T any]() *DLDeque[T] {
	return &DLDeque[T]{
		data: dllist.NewDLList[T](),
	}
}

func (dld *DLDeque[T]) Size() int {
	return dld.data.Size()
}

func (dld *DLDeque[T]) AddFirst(x T) {
	dld.data.Add(0, x)
}

func (dld *DLDeque[T]) AddLast(x T) {
	dld.data.Add(dld.Size(), x)
}

func (dld *DLDeque[T]) RemoveFirst() (T, error) {
	var result, err = dld.data.Remove(0)
	if err != nil {
		return *new(T), fmt.Errorf("DLDeque is empty")
	}
	return result, err
}

func (dld *DLDeque[T]) RemoveLast() (T, error) {
	var result, err = dld.data.Remove(dld.Size() - 1)
	if err != nil {
		return *new(T), fmt.Errorf("DLDeque is empty")
	}
	return result, err
}
