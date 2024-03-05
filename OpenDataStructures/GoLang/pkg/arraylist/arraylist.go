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
		al.resize()
	}
	if i < al.size/2 {
		al.head = utils.Mod(al.head-1, len(al.data))
		for j := 0; j < i; j++ {
			to := utils.Mod(al.head+j, len(al.data))
			from := utils.Mod(al.head+j+1, len(al.data))
			al.data[to] = al.data[from]
		}
	} else {
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
	if i < 0 || i >= al.size {
		return nullValue, fmt.Errorf("Index out of range: %d", i)
	}
	result := al.data[utils.Mod(i+al.head, len(al.data))]
	if i < al.size/2 {
		for j := i; j > 0; j-- {
			from := utils.Mod(j+al.head-1, len(al.data))
			to := utils.Mod(j+al.head, len(al.data))
			al.data[to] = al.data[from]
		}
		al.data[al.head] = nullValue
		al.head = utils.Mod(al.head+1, len(al.data))
	} else {
		for j := i; j < al.size-1; j++ {
			from := utils.Mod(i+al.head+1, len(al.data))
			to := utils.Mod(j+al.head, len(al.data))
			al.data[to] = al.data[from]
		}
		al.data[utils.Mod(al.head+al.size-1, len(al.data))] = nullValue
	}
	al.size--
	if len(al.data) >= al.size*3 {
		al.resize()
	}
	return result, nil
}

func (al *ArrayList[T]) realign(newData []T) {
	lengthOfSource := len(al.data)
	if al.head+al.size > lengthOfSource {
		copy(newData[0:lengthOfSource-al.head], al.data[al.head:lengthOfSource])
		copy(newData[lengthOfSource-al.head:al.size], al.data[0:utils.Mod(al.head+al.size, lengthOfSource)])
	} else {
		copy(newData[0:al.size], al.data[al.head:al.head+al.size])
	}
	al.data = newData
	al.head = 0
}

func (al *ArrayList[T]) resize() {
	newSize := 1
	if al.size*2 > newSize {
		newSize = al.size * 2
	}
	newData := make([]T, newSize)
	al.realign(newData)
}
