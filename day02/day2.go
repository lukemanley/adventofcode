package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var dimensions [][]int
	for scanner.Scan() {
		strs := strings.Split(scanner.Text(), "x")
		ints := make([]int, 3, 3)

		for i, s := range strs {
			x, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal(err)
			}
			ints[i] = x
		}
		sort.Ints(ints)
		dimensions = append(dimensions, ints)
	}

	fmt.Printf("Solution 1: %d\n", computePaper(dimensions))
	fmt.Printf("Solution 2: %d\n", computeRibbon(dimensions))
}

func computePaper(dimensions [][]int) int {
	paper := 0
	for _, box := range dimensions {
		d1, d2, d3 := box[0], box[1], box[2]
		s1 := d1 * d2
		s2 := d1 * d3
		s3 := d2 * d3
		paper += d1 + 2*s1 + 2*s2 + 2*s3
	}
	return paper
}

func computeRibbon(dimensions [][]int) int {
	ribbon := 0
	for _, box := range dimensions {
		d1, d2, d3 := box[0], box[1], box[2]
		ribbon += 2*(d1+d2) + d1*d2*d3
	}
	return ribbon
}
