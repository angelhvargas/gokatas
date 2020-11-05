package main

import (
	"testing"
)

func TestWordWrap(t *testing.T) {
	var tests = []struct {
		words string
		lines_expected int

	}{
		{"this", 1},
		{"this do that", 3},
		{"if this is longer than 3 then 8", 8}
	}

	for _, test := range tests {
		if output := WordWrap(test.words); output != test.lines_expected  {
			t.Error("Test Failed: {} inputted, {} expected, received {}", test.input, test.lines_expected, output)
		}
	}
}
