package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntComparer(t *testing.T) {
	a := 0
	b := 1
	c := 1

	assert.Equal(t, -1, IntComparer(a, b))
	assert.Equal(t, 1, IntComparer(b, a))
	assert.Equal(t, 0, IntComparer(b, c))
}
