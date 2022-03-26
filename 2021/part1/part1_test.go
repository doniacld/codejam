package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetTreats(t *testing.T) {
	sizes := []int{10, 20, 10, 25}
	out := getMin(sizes)
	assert.Equal(t, 7, out)

	sizes = []int{7, 7, 7, 7, 7}
	out = getMin(sizes)
	assert.Equal(t, 5, out)


	sizes = []int{100,1}
	out = getMin(sizes)
	assert.Equal(t, 3, out)
}
