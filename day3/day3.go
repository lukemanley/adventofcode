package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	f, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err.Error())
	}

	var x, y int
	xy := make(map[[2]int]int)
	xy[[2]int{x, y}]++

	for _, r := range f {
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
	fmt.Println("Single:", len(xy))

	var x1, y1, x2, y2 int
	xy = make(map[[2]int]int)
	xy[[2]int{x1, y1}]++
	xy[[2]int{x2, y2}]++

	for i, r := range f {
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
	fmt.Println("Double:", len(xy))

}
