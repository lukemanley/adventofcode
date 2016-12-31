package main

import (
	"bytes"
	"fmt"
)

func main() {

	input := "1113222113"

	for i := 0; i < 50; i++ {
		if i == 40 {
			fmt.Println("Solution 1:", len(input))
		}
		input = transform(input)
	}

	fmt.Println("Solution 2:", len(input))
}

func transform(input string) string {

	var output bytes.Buffer
	n := 0
	var char rune

	for i, c := range input {

		if i > 0 && c != char {
			output.WriteRune(rune(48 + n))
			output.WriteRune(char)
			n = 0
		}

		n++
		char = c

	}

	output.WriteRune(rune(48 + n))
	output.WriteRune(char)

	return output.String()
}
