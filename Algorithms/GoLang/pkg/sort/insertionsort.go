package sort

func InsertionSort[T any](comparer func(T,T) int, values []T) []T {
	//check each element and place in the first section of the slice where it belongs

	for i := 0; i < len(values); i++ {
		for j := i; j > 0; j-- {
			tmp := values[j]
			if comparer(values[j-1], tmp) > 0 {
				values[j], values[j-1] = values[j-1], values[j]
			}else {
				break
			}
		}
	}
	return values
}
