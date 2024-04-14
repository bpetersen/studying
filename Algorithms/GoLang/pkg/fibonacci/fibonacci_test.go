package fibonacci

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func FibTest(t *testing.T, function func(i int) int) {
	actual := function(0)
	assert.Equal(t, 1, actual)
	actual = function(1)
	assert.Equal(t, 1, actual)
	actual = function(2)
	assert.Equal(t, 2, actual)
	actual = function(3)
	assert.Equal(t, 3, actual)
	actual = function(4)
	assert.Equal(t, 5, actual)
	actual = function(7)
	assert.Equal(t, 21, actual)
}

func TestFibonacciRecursive(t *testing.T) {
	function := FibonacciRecursive
	FibTest(t, function)
}

func TestFibonacciStack(t *testing.T) {
	function := FibonacciStack
	FibTest(t, function)
}

func TestFibonacciMemoized(t *testing.T) {
	function := FibonacciMemoized
	FibTest(t, function)
}

func TestTimes(t *testing.T) {
	start := time.Now()
	_ = FibonacciRecursive(40)
	recursiveDuration := time.Since(start)

	start = time.Now()
	_ = FibonacciStack(40)
	stackDuration := time.Since(start)

	start = time.Now()
	_ = FibonacciMemoized(40)
	memoDuration := time.Since(start)

	assert.Equal(t, 0, recursiveDuration)
	assert.Equal(t, 0, stackDuration)
	assert.Equal(t, 0, memoDuration)
}
