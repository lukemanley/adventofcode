package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var m [1000000]int

	f, err := os.Open("input.txt")
	if err != nil {
		panic(err.Error())
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()

		var coords = make([]int, 4)
		pieces1 := strings.Split(line, " ")

		idx := 0
		for _, s := range pieces1 {
			if strings.Contains(s, ",") {
				s2 := strings.Split(s, ",")
				for _, c := range s2 {
					i, _ := strconv.Atoi(c)
					coords[idx] = i
					idx++
				}
			}
		}

		x1 := coords[0]
		x2 := coords[2]
		if x1 > x2 {
			x1, x2 = x2, x1
		}

		y1 := coords[1]
		y2 := coords[3]
		if y1 > y2 {
			y1, y2 = y2, y1
		}

		switch {
		case strings.HasPrefix(line, "turn on "):
			for x := x1; x <= x2; x++ {
				for y := y1; y <= y2; y++ {
					m[1000*x+y]++
				}
			}

		case strings.HasPrefix(line, "turn off "):
			for x := x1; x <= x2; x++ {
				for y := y1; y <= y2; y++ {
					v := m[1000*x+y]
					if v > 0 {
						m[1000*x+y]--
					}
				}
			}

		case strings.HasPrefix(line, "toggle "):
			for x := x1; x <= x2; x++ {
				for y := y1; y <= y2; y++ {
					m[1000*x+y]++
					m[1000*x+y]++
				}
			}
		}
	}

	sum := 0
	for _, on := range m {
		sum += on
	}
	fmt.Println(sum)
}
