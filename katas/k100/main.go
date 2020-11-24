package main

import "strings"

// WordWrap (text string) (string, int): receives an string
// and returns the words wrapped to 1 per line
func WordWrap(text string) (string, int) {
	var textWrapped []string
	var textOutput string
	var linesCount int

	textWrapped = strings.Fields(text)
	textOutput = strings.Join(textWrapped, "\n")

	linesCount = len(textWrapped)

	return textOutput, linesCount
}

func main() {

}
