package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type instruct struct {
	instruct       string
	x1, y1, x2, y2 int
}

var re = regexp.MustCompile(`(\bturn on\b|\bturn off\b|\btoggle\b) (\d+),(\d+) through (\d+),(\d+)`)

func main() {

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	var instructs []instruct

	for s.Scan() {
		line := s.Text()

		pieces := re.FindStringSubmatch(line)

		x1, _ := strconv.Atoi(pieces[2])
		y1, _ := strconv.Atoi(pieces[3])
		x2, _ := strconv.Atoi(pieces[4])
		y2, _ := strconv.Atoi(pieces[5])

		if x1 > x2 {
			x1, x2 = x2, x1
		}
		if y1 > y2 {
			y1, y2 = y2, y1
		}

		ins := instruct{
			pieces[1],
			x1,
			y1,
			x2,
			y2,
		}
		instructs = append(instructs, ins)
	}

	fmt.Printf("Solution 1: %d\n", p1(instructs))
	fmt.Printf("Solution 2: %d\n", p2(instructs))
}

func p1(instructs []instruct) int {
	var m [1000000]int
	for _, ins := range instructs {
		switch ins.instruct {
		case "turn on":
			for x := ins.x1; x <= ins.x2; x++ {
				for y := ins.y1; y <= ins.y2; y++ {
					m[1000*x+y] = 1
				}
			}

		case "turn off":
			for x := ins.x1; x <= ins.x2; x++ {
				for y := ins.y1; y <= ins.y2; y++ {
					m[1000*x+y] = 0
				}
			}

		case "toggle":
			for x := ins.x1; x <= ins.x2; x++ {
				for y := ins.y1; y <= ins.y2; y++ {
					m[1000*x+y]--
					m[1000*x+y] *= -1
				}
			}
		}
	}
	sum := 0
	for _, on := range m {
		sum += on
	}
	return sum
}

func p2(instructs []instruct) int {
	var m [1000000]int
	for _, ins := range instructs {
		switch ins.instruct {
		case "turn on":
			for x := ins.x1; x <= ins.x2; x++ {
				for y := ins.y1; y <= ins.y2; y++ {
					m[1000*x+y]++
				}
			}

		case "turn off":
			for x := ins.x1; x <= ins.x2; x++ {
				for y := ins.y1; y <= ins.y2; y++ {
					if m[1000*x+y] > 0 {
						m[1000*x+y]--
					}
				}
			}

		case "toggle":
			for x := ins.x1; x <= ins.x2; x++ {
				for y := ins.y1; y <= ins.y2; y++ {
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
	return sum
}
