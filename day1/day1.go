package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Printf("problem 1: %d\n", p1())
	fmt.Printf("problem 2: %d\n", p2())
}

func p1() int {

	f, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err.Error())
	}

	floor := 0

	for _, v := range string(f) {
		switch v {
		case '(':
			floor++
		case ')':
			floor--
		}
	}

	return floor
}

func p2() int {

	f, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err.Error())
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

	return basement
}
