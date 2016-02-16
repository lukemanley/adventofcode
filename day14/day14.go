package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var re = regexp.MustCompile(`([A-Za-z]+) can fly ([0-9]+) km/s for ([0-9]+) seconds, but then must rest for ([0-9]+) seconds.`)

func main() {
	p1()
	p2()
}

func p1() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)

	max := 0

	for s.Scan() {
		line := s.Text()
		m := re.FindStringSubmatch(line)

		speed, err := strconv.Atoi(m[2])
		if err != nil {
			panic(err)
		}
		duration, err := strconv.Atoi(m[3])
		if err != nil {
			panic(err)
		}
		rest, err := strconv.Atoi(m[4])
		if err != nil {
			panic(err)
		}

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
	fmt.Println("Solution 1:", max)
}

func p2() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)

	data := make(map[string][]int)

	for s.Scan() {
		line := s.Text()
		m := re.FindStringSubmatch(line)

		speed, err := strconv.Atoi(m[2])
		if err != nil {
			panic(err)
		}
		duration, err := strconv.Atoi(m[3])
		if err != nil {
			panic(err)
		}
		rest, err := strconv.Atoi(m[4])
		if err != nil {
			panic(err)
		}

		data[m[1]] = []int{speed, duration, rest}
	}

	dist := make(map[string]int)
	points := make(map[string]int)

	for i := 0; i < 2503; i++ {
		maxdist := 0
		for name, stats := range data {
			speed := stats[0]
			duration := stats[1]
			rest := stats[2]
			var d int = i / (duration + rest)
			var mod int = i - d*(duration+rest)
			if mod < duration {
				dist[name] += speed
			}
			if dist[name] > maxdist {
				maxdist = dist[name]
			}
		}
		for name, d := range dist {
			if d == maxdist {
				points[name]++
			}
		}
	}

	max := 0
	for _, pts := range points {
		if pts > max {
			max = pts
		}
	}
	fmt.Println("Solution 2:", max)
}
