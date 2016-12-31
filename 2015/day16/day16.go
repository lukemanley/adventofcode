package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var re = regexp.MustCompile(`Sue (\d+): (\w+): (\d+), (\w+): (\d+), (\w+): (\d+)`)

var m = map[string]int{
	"children":    3,
	"cats":        7,
	"samoyeds":    2,
	"pomeranians": 3,
	"akitas":      0,
	"vizslas":     0,
	"goldfish":    5,
	"trees":       3,
	"cars":        2,
	"perfumes":    1,
}

type sue struct {
	num   int
	items []string
	vals  []int
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	var sues []sue

	for s.Scan() {
		sli := re.FindStringSubmatch(s.Text())

		num, err := strconv.Atoi(sli[1])
		if err != nil {
			log.Fatal(err)
		}

		sue2 := sue{num, nil, nil}

		for i := 2; i < len(sli); i += 2 {
			item := sli[i]
			n, err := strconv.Atoi(sli[i+1])
			if err != nil {
				log.Fatal(err)
			}
			sue2.items = append(sue2.items, item)
			sue2.vals = append(sue2.vals, n)
		}
		sues = append(sues, sue2)
	}

	fmt.Println("Solution 1:", p1(sues))
	fmt.Println("Solution 2:", p2(sues))
}

func p1(sues []sue) int {
	for _, s := range sues {
		for i, item := range s.items {
			if m[item] != s.vals[i] {
				break
			}
			if i == len(s.items)-1 {
				return s.num
			}
		}
	}
	return -1
}

func p2(sues []sue) int {
	for _, s := range sues {
	loop:
		for i, item := range s.items {
			n := s.vals[i]
			switch item {
			case "cats", "trees":
				if m[item] >= n {
					break loop
				}
			case "pomeranians", "goldfish":
				if m[item] <= n {
					break loop
				}
			default:
				if m[item] != n {
					break loop
				}
			}
			if i == len(s.items)-1 {
				return s.num
			}
		}
	}
	return -1
}
