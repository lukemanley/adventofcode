package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var re = regexp.MustCompile(`(\w+) would (\w+) (\d+) happiness units by sitting next to (\w+).`)

type seat struct {
	a string
	b string
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var people []string
	happiness := make(map[seat]int)

	s := bufio.NewScanner(f)

	for s.Scan() {
		match := re.FindStringSubmatch(s.Text())

		v, err := strconv.Atoi(match[3])
		if err != nil {
			log.Fatal(err)
		}
		if match[2] == "lose" {
			v = -v
		}

		if !containsString(people, match[1]) {
			people = append(people, match[1])
		}
		happiness[seat{match[1], match[4]}] = v
	}

	seated := make([]int, len(people), len(people))
	fmt.Println("Solution 1:", best(happiness, people, seated, 0, 0))
	fmt.Println("Solution 2:", best2(happiness, people, seated, 0, 0))

}

func best(happiness map[seat]int, people []string, seated []int, i, max int) int {

	if i == len(people) {
		h := 0
		b := people[seated[i-1]]
		for _, j := range seated {
			a := people[j]
			h += happiness[seat{a, b}]
			h += happiness[seat{b, a}]
			b = a
		}
		if h > max {
			max = h
		}
		return max
	}

	for j := 0; j < len(people); j++ {
		if containsInt(seated[:i], j) {
			continue
		}
		seated[i] = j
		max = best(happiness, people, seated, i+1, max)
	}
	return max
}

func best2(happiness map[seat]int, people []string, seated []int, i, max int) int {

	if i == len(people) {
		h := 0
		min := 0
		minset := false
		b := people[seated[i-1]]
		for _, j := range seated {
			a := people[j]
			h2 := happiness[seat{a, b}] + happiness[seat{b, a}]
			if !minset || h2 < min {
				min = h2
				minset = true
			}
			h += h2
			b = a
		}
		h -= min
		if h > max {
			max = h
		}
		return max
	}

	for j := 0; j < len(people); j++ {
		if containsInt(seated[:i], j) {
			continue
		}
		seated[i] = j
		max = best2(happiness, people, seated, i+1, max)
	}
	return max
}

func containsInt(sli []int, i int) bool {
	for _, j := range sli {
		if i == j {
			return true
		}
	}
	return false
}

func containsString(sli []string, s string) bool {
	for _, s2 := range sli {
		if s == s2 {
			return true
		}
	}
	return false
}
