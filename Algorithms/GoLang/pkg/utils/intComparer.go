package utils

// should be a - b.  So if B is larger, you'll have a negative number
func IntComparer(a int, b int) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	}
	return 0
}
