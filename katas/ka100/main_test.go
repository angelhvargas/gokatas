package main

import (
	"testing"
)

func TestWordWrap(t *testing.T) {
	var tests = []struct {
		words         string
		linesExpected int
	}{
		{"foo", 1},
		{"foo do bar", 3},
		{"if this is longer than 3 then 8", 8},
		{"if this is  double spaces we want  7", 8},
		{"", 0},
		{"  test  ", 1},
	}

	for _, test := range tests {
		if _, output := WordWrap(test.words); output != test.linesExpected {
			t.Error("Test Failed: {} inputted, {} expected, received {}", test.words, test.linesExpected, output)
		}
	}
}
