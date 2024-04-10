package binarysearch

//learned about the bitshift from: https://research.google/blog/extra-extra-read-all-about-it-nearly-all-binary-searches-and-mergesorts-are-broken/

// return the index for the item in the slice
// if not found, return the index for where the item would go
func BinarySearch[T any](comparer func(T, T) int, items []T, item T) (int, bool) {
	left := 0
	right := len(items) - 1
	for left <= right {
		mid := (left + right) >> 1
		midValue := items[mid]
		if comparer(midValue, item) == 0 {
			return mid, true
		} else if comparer(midValue, item) < 0 {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	//return the index for where the item would go
	return left, false
}
