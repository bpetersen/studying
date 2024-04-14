package fibonacci

func FibonacciRecursive(i int) int {
	if i == 0 {
		return 1
	}
	if i == 1 {
		return 1
	}
	return FibonacciRecursive(i-1) + FibonacciRecursive(i-2)
}

func FibonacciMemoized(i int) int {
	if i < 2 {
		return 1
	}

	cache := make([]int, i+1)
	cache[0] = 1
	cache[1] = 1
	return fibonacciMemoized(i, cache)
}

func fibonacciMemoized(i int, cache []int) int {
	if cache[i] != 0 {
		return cache[i]
	}
	result := fibonacciMemoized(i-1, cache) + fibonacciMemoized(i-2, cache)
	cache[i] = result
	return result
}

func FibonacciStack(i int) int {
	stack := []int{}
	stack = Push[int](stack, i)

	result := 0
	currentValue := 0

	for len(stack) > 0 {
		stack, currentValue = Pop[int](stack)
		if currentValue == 0 || currentValue == 1 {
			result += 1
		} else {
			stack = append(stack, currentValue-1)
			stack = append(stack, currentValue-2)
		}
	}
	return result
}

func Push[T any](stack []T, value T) []T {
	stack = append(stack, value)

	return stack
}

func Pop[T any](stack []T) ([]T, T) {
	if len(stack) == 0 {
		panic("pop from empty stack.")
	}
	value := stack[len(stack)-1]
	stack = stack[:len(stack)-1]

	return stack, value
}
