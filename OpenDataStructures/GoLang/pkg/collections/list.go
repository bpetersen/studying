package collections

type List[T any] interface {
	Size() int
	Get(i int) (T, error)      // Get the item at position i
	Set(i int, x T) (T, error) // Set the value of the array at position i equal to x and returns the value being replaced.
	Add(i int, x T) error      // Add the value x at position i and shifting others to make room
	Remove(i int) (T, error)   // Remove the value at position i and shift others to fill in the empty slot
}
