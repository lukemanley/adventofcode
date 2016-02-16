package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Solution 1: %d\n", p1(b))
	fmt.Printf("Solution 2: %d\n", p2(b))
}

func p1(b []byte) int {

	floor := 0

	for _, r := range string(b) {
		switch r {
		case '(':
			floor++
		case ')':
			floor--
		}
	}

	return floor
}

func p2(b []byte) int {

	floor := 0
	basement := 0

	for i, r := range string(b) {
		switch r {
		case '(':
			floor++
		case ')':
			floor--
		}
		if basement == 0 && floor < 0 {
			basement = i + 1
		}
	}

	return basement
}
