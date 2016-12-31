package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type tri [3]int

func main() {
	tris := parse("input.txt")

	// part 1
	fmt.Println(p1(tris))

	// part 2
	fmt.Println(p2(tris))
}

func parse(filename string) []tri {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	s := strings.TrimSpace(string(b))
	lines := strings.Split(s, "\n")
	tris := []tri{}
	for _, line := range lines {
		var a, b, c int
		_, err := fmt.Sscanf(line, "%d %d %d", &a, &b, &c)
		if err != nil {
			log.Fatal(err)
		}
		tris = append(tris, tri{a, b, c})
	}
	return tris
}

func p1(tris []tri) int {
	var n int
	for _, tri := range tris {
		var sum int
		var max int
		for _, v := range tri {
			sum += v
			if v > max {
				max = v
			}
		}
		if max*2 < sum {
			n++
		}
	}
	return n
}

func p2(tris []tri) int {
	var n int
	for i := 0; i < len(tris)/3; i++ {
		for j := 0; j < 3; j++ {
			var sum int
			var max int
			for k := 0; k < 3; k++ {
				v := tris[i*3+k][j]
				sum += v
				if v > max {
					max = v
				}
			}
			if max*2 < sum {
				n++
			}
		}
	}
	return n
}
