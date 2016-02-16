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

func main() {
	p1()
	p2()
}

func p1() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)

	for s.Scan() {
		line := s.Text()
		sli := re.FindStringSubmatch(line)
		for i := 0; i < 3; i++ {
			item := sli[2+i*2]
			n, err := strconv.Atoi(sli[3+i*2])
			if err != nil {
				log.Fatal(err)
			}
			if m[item] != n {
				break
			}
			if i == 2 {
				fmt.Println("Solution 1:", sli[1])
			}
		}
	}
}

func p2() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)

	for s.Scan() {
		line := s.Text()
		sli := re.FindStringSubmatch(line)

	loop:
		for i := 0; i < 3; i++ {
			item := sli[2+i*2]
			n, err := strconv.Atoi(sli[3+i*2])
			if err != nil {
				log.Fatal(err)
			}
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
			if i == 2 {
				fmt.Println("Solution 2:", sli[1])
			}
		}
	}
}
