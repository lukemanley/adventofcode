package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	f, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic("Error reading file")
	}

	floor := 0
	basement := 0

	for i, v := range string(f) {
		switch v {
		case '(':
			floor++
		case ')':
			floor--
		}
		if basement == 0 && floor < 0 {
			basement = i + 1
		}
	}

	fmt.Println("final floor:", floor)
	fmt.Println("basement index:", basement)
}
