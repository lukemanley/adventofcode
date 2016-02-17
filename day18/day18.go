package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	var lights [10000]int
	i := 0

	for s.Scan() {
		for _, r := range s.Text() {
			switch r {
			case '#':
				lights[i] = 1
			case '.':
				// off
			default:
				log.Fatal("unrecognized input: " + string(r))
			}
			i++
		}
	}

	fmt.Println("Solution 1:", p1(lights))
	fmt.Println("Solution 2:", p2(lights))
}

func p1(lights [10000]int) int {
	for n := 0; n < 100; n++ {
		lights = step(lights[:])
	}

	sum := 0
	for _, v := range lights {
		sum += v
	}
	return sum
}

func p2(lights [10000]int) int {
	for n := 0; n < 100; n++ {
		lights[0] = 1
		lights[99] = 1
		lights[9900] = 1
		lights[9999] = 1
		lights = step(lights[:])
		lights[0] = 1
		lights[99] = 1
		lights[9900] = 1
		lights[9999] = 1
	}

	sum := 0
	for _, v := range lights {
		sum += v
	}
	return sum
}

func step(lights []int) [10000]int {
	var new [10000]int

	for i, on := range lights {
		n := int(math.Mod(float64(i), 100))

		xmin, ymin := -1, -1
		xmax, ymax := 1, 1

		if n == 99 {
			xmax = 0
		}
		if n == 0 {
			xmin = 0
		}
		if i <= 99 {
			ymin = 0
		}
		if i >= 9900 {
			ymax = 0
		}

		neighborsum := 0

		for x := xmin; x <= xmax; x++ {
			for y := ymin; y <= ymax; y++ {
				if x == 0 && y == 0 {
					continue
				}
				neighborsum += lights[i+x+100*y]
			}
		}
		if on == 1 {
			if neighborsum == 2 || neighborsum == 3 {
				new[i] = 1
			}
		} else {
			if neighborsum == 3 {
				new[i] = 1
			}
		}
	}
	return new
}
