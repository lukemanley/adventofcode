package main

import (
	"fmt"
	"math"
)

const n = 33100000

func main() {
	fmt.Println("Solution 1:", f1())
	fmt.Println("Solution 2:", f2())
}

func f1() int {
	for x := 1; ; x++ {
		sum := 0
		sr := math.Sqrt(float64(x))
		for y := 1; y <= int(sr); y++ {
			if x%y == 0 {
				sum += 10 * y
				if float64(y) < sr {
					sum += 10 * x / y
				}
			}
		}
		if sum >= n {
			return x
		}
	}
	return -1
}

func f2() int {
	for x := 1; ; x++ {
		sum := 0
		sr := math.Sqrt(float64(x))
		if sr > 50 {
			sr = 50
		}
		for y := 1; y <= int(sr); y++ {
			if x%y == 0 {
				if y*50 <= x {
					sum += 11 * y
				}
				if float64(y) < sr {
					sum += 11 * x / y
				}
			}
		}
		if sum >= n {
			return x
		}
	}
	return -1
}
