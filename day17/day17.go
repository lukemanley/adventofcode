package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)

	var caps []int

	for s.Scan() {
		capacity, err := strconv.Atoi(s.Text())
		if err != nil {
			log.Fatal(err)
		}
		caps = append(caps, capacity)
	}

	counts := map[int]int{}

	n := fill(caps, 0, 0, 0, counts)

	min := 0
	for count, _ := range counts {
		if min == 0 {
			min = count
		} else if count < min {
			min = count
		}
	}

	fmt.Println("Solution 1:", n)
	fmt.Println("Solution 2:", counts[min])
}

func fill(caps []int, i int, total int, count int, counts map[int]int) int {

	if i == len(caps) {
		return 0
	}

	n := fill(caps, i+1, total, count, counts)

	total += caps[i]
	count++

	switch {
	case total < 150:
		n += fill(caps, i+1, total, count, counts)
	case total == 150:
		n += 1
		counts[count]++
	}
	return n
}
