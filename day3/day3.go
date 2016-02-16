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

	var x, y int
	xy := make(map[[2]int]int)
	xy[[2]int{x, y}]++

	for _, r := range b {
		switch r {
		case '^':
			y++
		case 'v':
			y--
		case '<':
			x--
		case '>':
			x++
		}
		xy[[2]int{x, y}]++
	}
	return len(xy)
}

func p2(b []byte) int {

	var x1, y1, x2, y2 int
	xy := make(map[[2]int]int)
	xy[[2]int{x1, y1}]++
	xy[[2]int{x2, y2}]++

	for i, r := range b {
		switch i%2 == 0 {
		case true:
			switch r {
			case '^':
				y1++
			case 'v':
				y1--
			case '<':
				x1--
			case '>':
				x1++
			}
			xy[[2]int{x1, y1}]++
		case false:
			switch r {
			case '^':
				y2++
			case 'v':
				y2--
			case '<':
				x2--
			case '>':
				x2++
			}
			xy[[2]int{x2, y2}]++

		}
	}
	return len(xy)
}
