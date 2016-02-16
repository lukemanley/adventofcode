package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type hop struct {
	from string
	to   string
}

var r = regexp.MustCompile(`^(\w+) to (\w+) = (\d+)$`)

func main() {

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	costs := make(map[hop]int)
	locationset := make(map[string]bool)

	for s.Scan() {
		m := r.FindStringSubmatch(s.Text())

		cost, err := strconv.Atoi(m[3])
		if err != nil {
			log.Fatal(err)
		}

		costs[hop{m[1], m[2]}] = cost
		locationset[m[1]] = true
		locationset[m[2]] = true
	}

	var routes [][]int

	route := make([]int, len(locationset), len(locationset))
	for i, _ := range route {
		route[i] = -1
	}

	routes = createRoutes(routes, route, len(locationset))

	var stops []string
	for k, _ := range locationset {
		stops = append(stops, k)
	}

	min := 0
	max := 0

	for _, r := range routes {

		cost := 0

		for i, j := range r {
			if i == 0 {
				continue
			}
			s1 := stops[r[i-1]]
			s2 := stops[j]

			c, ok := costs[hop{s1, s2}]
			if !ok {
				c, ok = costs[hop{s2, s1}]
				if !ok {
					log.Fatal("invalid hop")
				}
			}
			cost += c
		}
		if min == 0 || cost < min {
			min = cost
		}
		if cost > max {
			max = cost
		}
	}

	fmt.Println("Solution 1:", min)
	fmt.Println("Solution 2:", max)
}

func createRoutes(routes [][]int, route []int, n int) [][]int {
	if n == 0 {
		return append(routes, route)
	}
	for i := 0; i < len(route); i++ {
		if contains(route, i) {
			continue
		}

		r := make([]int, len(route), len(route))
		for idx, val := range route {
			r[idx] = val
		}
		r[n-1] = i

		routes = createRoutes(routes, r, n-1)
	}
	return routes
}

func contains(route []int, x int) bool {
	for _, y := range route {
		if x == y {
			return true
		}
	}
	return false
}
