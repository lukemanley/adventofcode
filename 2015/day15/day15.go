package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type ingredient struct {
	name       string
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	opt := make([]*ingredient, 0, 4)

	for s.Scan() {
		var i ingredient
		opt = append(opt, &i)
		fmt.Sscanf(s.Text(), "%s capacity %d, durability %d, flavor %d, texture %d, calories %d", &i.name, &i.capacity, &i.durability, &i.flavor, &i.texture, &i.calories)
	}

	maxscore := 0
	maxscore2 := 0
	for a := 0; a <= 100; a++ {
		for b := 0; a+b <= 100; b++ {
			for c := 0; a+b+c <= 100; c++ {
				d := 100 - a - b - c
				cap := a*opt[0].capacity + b*opt[1].capacity + c*opt[2].capacity + d*opt[3].capacity
				dur := a*opt[0].durability + b*opt[1].durability + c*opt[2].durability + d*opt[3].durability
				fla := a*opt[0].flavor + b*opt[1].flavor + c*opt[2].flavor + d*opt[3].flavor
				tex := a*opt[0].texture + b*opt[1].texture + c*opt[2].texture + d*opt[3].texture

				if cap <= 0 || dur <= 0 || fla <= 0 || tex <= 0 {
					continue
				}

				score := cap * dur * fla * tex
				if score > maxscore {
					maxscore = score
				}

				cal := a*opt[0].calories + b*opt[1].calories + c*opt[2].calories + d*opt[3].calories
				if cal == 500 && score > maxscore2 {
					maxscore2 = score
				}
			}
		}
	}
	fmt.Println("Solution 1:", maxscore)
	fmt.Println("Solution 2:", maxscore2)
}
