package arraylist

import (
	"fmt"
	"open_data_structures/pkg/collections"
	"open_data_structures/pkg/utils"
)

var _ collections.List[int] = &ArrayList[int]{}

const DEFAULT_INITIAL_CAPACITY = 1

type ArrayList[T any] struct {
	data []T //The backing slice that contains the data
	size int //The number of elements in the list
	head int //The index of the first element in the backing slice
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

func (al *ArrayList[T]) Size() int {
	return al.size
}

func (al *ArrayList[T]) Get(i int) (T, error) {
	var nullValue T
	if i < 0 || i >= al.size {
		return nullValue, fmt.Errorf("Index out of range: %d", i)
	}
	return al.data[utils.Mod(i+al.head, len(al.data))], nil
}

func (al *ArrayList[T]) Set(i int, x T) (T, error) {
	var nullValue T
	if i < 0 || i >= al.size {
		return nullValue, fmt.Errorf("Index out of range: %d", i)
	}
	index := utils.Mod(i+al.head, len(al.data))
	result := al.data[index]
	al.data[index] = x
	return result, nil
}

func (al *ArrayList[T]) Add(i int, x T) error {
	if i < 0 || i > al.size {
		return fmt.Errorf("Index out of range: %d", i)
	}
	if len(al.data) == al.size {
		newData := make([]T, len(al.data)*2)
		copy(newData, al.data)
		al.data = newData
	}
	if i < al.size/2 {
		al.head = utils.Mod(al.head-1, len(al.data))
		for j := 0; j < i; j++ {
			to := utils.Mod(al.head+j, len(al.data))
			from := utils.Mod(al.head+j+1, len(al.data))
			al.data[to] = al.data[from]
		}
	} else {
		fmt.Println(al.data)
		for j := al.size; j > i; j-- {
			to := utils.Mod(al.head+j, len(al.data))
			from := utils.Mod(al.head+j-1, len(al.data))
			al.data[to] = al.data[from]
		}
	}
	al.data[utils.Mod(al.head+i, len(al.data))] = x
	al.size++
	return nil
}

func (al *ArrayList[T]) Remove(i int) (T, error) {
	var nullValue T
	return nullValue, nil
}
