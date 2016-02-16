package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

var re1 = regexp.MustCompile(`\\\"`)
var re2 = regexp.MustCompile(`\\\\`)
var re3 = regexp.MustCompile(`\\x..`)

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

	fmt.Printf("Solution 1: %d\n", p1(lines))
	fmt.Printf("Solution 2: %d\n", p2(lines))
}

func p1(lines []string) int {

	ncode := 0
	nstr := 0

	for _, line := range lines {
		ncode += len(line)

		line = re1.ReplaceAllString(line, "_")
		line = re2.ReplaceAllString(line, "_")
		line = re3.ReplaceAllString(line, "_")

		nstr += len(line)
		nstr -= 2 // line quotes
	}
	return ncode - nstr
}

func p2(lines []string) int {

	ncode := 0
	nstr := 0

	for _, line := range lines {
		ncode += len(line)

		line = re1.ReplaceAllString(line, "____")
		line = re2.ReplaceAllString(line, "____")
		line = re3.ReplaceAllString(line, "_____")

		nstr += len(line)
		nstr += 4 // line quotes
	}

	return nstr - ncode
}
