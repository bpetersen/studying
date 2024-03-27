package utils

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHeight0(t *testing.T) {
	binaryString := "100010"
	num, _ := strconv.ParseInt(binaryString, 2, 64)
	height := GetRandomHeight(int(num))
	assert.Equal(t, 0, height)
}

func TestGetHeight1(t *testing.T) {
	height := GetRandomHeight(1)
	assert.Equal(t, 1, height)
}

func TestGetHeight7(t *testing.T) {
	height := GetRandomHeight(7)
	assert.Equal(t, 3, height)
}
