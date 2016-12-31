package main

import "fmt"

func main() {

	row := 2947
	col := 3029

	cols := col + row - 1

	cells := 0
	for n := 1; n < cols; n++ {
		cells += n
	}

	cells += col

	v := 20151125
	for n := 1; n < cells; n++ {
		v *= 252533
		v = v % 33554393
	}
	fmt.Println("Solution 1:", v)

}
