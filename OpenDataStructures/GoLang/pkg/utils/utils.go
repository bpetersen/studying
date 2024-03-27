package utils

func Mod(a int, b int) int {
	return ((a % b) + b) % b
}

type Pair[T any, U any] struct {
	First  T
	Second U
}

// returns the number of "heads in a row". Used for determining the height of an item in a skiplist.
func GetRandomHeight(num int) int {
	count := 0
	for num > 0 {
		if num&1 == 1 {
			count++
		} else {
			break // Stop counting when the first 0 is encountered from the right
		}
		num >>= 1
	}

	return count
}
