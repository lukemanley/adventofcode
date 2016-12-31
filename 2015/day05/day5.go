package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	var lines []string
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	fmt.Println("Solution 1:", p1(lines))
	fmt.Println("Solution 2:", p2(lines))
}

func p1(lines []string) int {
	count := 0
	for _, line := range lines {
		if !repeats(line) {
			continue
		}
		if !strings.Contains(line, "ab") && !strings.Contains(line, "cd") && !strings.Contains(line, "pq") && !strings.Contains(line, "xy") {
			vowels := 0
			for _, r := range "aeiou" {
				vowels += strings.Count(line, string(r))
			}
			if vowels > 2 {
				count++
			}
		}
	}
	return count
}

func p2(lines []string) int {
	count := 0
	for _, line := range lines {
		pair := false
		repeat := false
		for i, r := range line {
			if i > 0 && !pair {
				s := strings.Split(line, fmt.Sprintf("%c%c", line[i-1], r))
				if len(s) > 2 {
					pair = true
				}
			}
			if i > 1 && !repeat {
				if r == rune(line[i-2]) {
					repeat = true
				}
			}
		}
		if pair && repeat {
			count++
		}
	}
	return count
}

func repeats(s string) bool {
	for i, r := range s[1:] {
		if r == rune(s[i]) {
			return true
		}
	}
	return false
}
