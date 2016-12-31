package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type pos struct {
	x, y int
}

func main() {
	moves := parse("input.txt")

	// part 1
	fmt.Println(p1(moves))

	// part 2
	fmt.Println(p2(moves))
}

func parse(filename string) []string {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	s := strings.TrimSpace(string(b))
	return strings.Split(s, "\n")
}

func p1(moves []string) string {
	keypad := [][]string{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	}

	p := pos{1, 1}
	code := ""
	for _, m := range moves {
		for _, r := range m {
			switch r {
			case 'U':
				if p.y > 0 {
					p.y--
				}
			case 'D':
				if p.y < 2 {
					p.y++
				}
			case 'R':
				if p.x < 2 {
					p.x++
				}
			case 'L':
				if p.x > 0 {
					p.x--
				}
			}
		}
		code += keypad[p.y][p.x]
	}
	return code
}

func p2(moves []string) string {
	keypad := [][]string{
		{"", "", "1", "", ""},
		{"", "2", "3", "4", ""},
		{"5", "6", "7", "8", "9"},
		{"", "A", "B", "C", ""},
		{"", "", "D", "", ""},
	}

	p := pos{3, 0}
	code := ""
	for _, m := range moves {
		for _, r := range m {
			switch r {
			case 'U':
				if p.y > 0 && keypad[p.y-1][p.x] != "" {
					p.y--
				}
			case 'D':
				if p.y < 4 && keypad[p.y+1][p.x] != "" {
					p.y++
				}
			case 'R':
				if p.x < 4 && keypad[p.y][p.x+1] != "" {
					p.x++
				}
			case 'L':
				if p.x > 0 && keypad[p.y][p.x-1] != "" {
					p.x--
				}
			}
		}
		code += keypad[p.y][p.x]
	}
	return code
}
