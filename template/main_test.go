package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_run(t *testing.T) {
	tt := map[string]struct {
		input, output string
	}{
		"ts1": {
			"testdata/ts1_input.txt",
			"testdata/ts1_output.txt",
		},
		"ts2": {
			"testdata/ts2_input.txt",
			"testdata/ts2_output.txt",
		},
		"ts3": {
			"testdata/ts3_input.txt",
			"testdata/ts3_output.txt",
		},
		"mine": {
			"testdata/mine_in.txt",
			"testdata/mine_out.txt",
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			f, _ := os.Open(tc.input)
			mi := MyInput{rdr: f}

			o, _ := os.Open(tc.output)
			mo := MyInput{rdr: o}
			run(mi, func(s string) {
				expected := mo.readLine()
				assert.Equal(t, expected, s)
			})
		})
	}
}
