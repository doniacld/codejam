package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetOccurrence(t *testing.T) {
	tt := []struct {
		description    string
		intput         string
		expectedOutput int
	}{
		{"IO appears twice", "IiOioIoO", 2},
		{"IO appears once", "IiOOIo", 1},
		{"IO does not appear", "IoiOiO", 0},
		{"IO does not appear", "io", 0},
		{"IO does not for times", "IIIIOOOO", 4},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			out := getOccurrence(tc.intput)
			assert.Equal(t, tc.expectedOutput, out)
		})
	}
}
