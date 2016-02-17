package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var re = regexp.MustCompile(`\w+ can fly (\d+) km/s for (\d+) seconds, but then must rest for (\d+) seconds.`)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	var stats [][]int

	for s.Scan() {
		m := re.FindStringSubmatch(s.Text())

		ints := make([]int, 3, 3)
		for i, str := range m[1:] {
			v, err := strconv.Atoi(str)
			if err != nil {
				log.Fatal(err)
			}
			ints[i] = v
		}
		stats = append(stats, ints)
	}

	fmt.Println("Solution 1:", p1(stats))
	fmt.Println("Solution 2:", p2(stats))
}

func p1(stats [][]int) int {
	max := 0

	for _, stats2 := range stats {

		speed := stats2[0]
		duration := stats2[1]
		rest := stats2[2]

		var d int = 2503 / (duration + rest)
		var mod int = 2503 - d*(duration+rest)

		dist := d * speed * duration

		if mod >= duration {
			dist += speed * duration
		} else {
			dist += speed * mod
		}

		if dist > max {
			max = dist
		}
	}
	return max
}

func p2(stats [][]int) int {

	dist := make(map[int]int)
	points := make(map[int]int)

	for i := 0; i < 2503; i++ {
		maxdist := 0
		for idx, stats2 := range stats {
			speed := stats2[0]
			duration := stats2[1]
			rest := stats2[2]
			var d int = i / (duration + rest)
			var mod int = i - d*(duration+rest)
			if mod < duration {
				dist[idx] += speed
			}
			if dist[idx] > maxdist {
				maxdist = dist[idx]
			}
		}
		for idx, d := range dist {
			if d == maxdist {
				points[idx]++
			}
		}
	}

	max := 0
	for _, pts := range points {
		if pts > max {
			max = pts
		}
	}
	return max
}
