package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)

	scanner.Split(bufio.ScanLines)

	paper := 0
	ribbon := 0

	for scanner.Scan() {
		line := scanner.Text()
		a := strings.Split(line, "x")

		if len(a) != 3 {
			panic("unexpected line format: " + line)
		}

		a2 := make([]int, 3, 3)

		for i, s := range a {

			x, err := strconv.Atoi(s)
			if err != nil {
				panic(err.Error())
			}
			a2[i] = x
		}

		sort.Ints(a2)

		d1, d2, d3 := a2[0], a2[1], a2[2]

		s1 := d1 * d2
		s2 := d1 * d3
		s3 := d2 * d3

		paper += d1 + 2*s1 + 2*s2 + 2*s3
		ribbon += 2*(d1+d2) + d1*d2*d3
	}

	fmt.Println("paper:", paper)
	fmt.Println("ribbon:", ribbon)
}
