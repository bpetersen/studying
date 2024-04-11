package sort

func BubbleSort[T comparable](comparer func(T, T) int, values []T) []T {
	for i := 1; i < len(values); i++ {
		for j := 0; j < len(values)-i; j++ {
			if comparer(values[j], values[j+1]) > 0 {
				values[j], values[j+1] = values[j+1], values[j]
			}
		}
	}
	return values
}
