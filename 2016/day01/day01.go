package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const (
	north = iota
	east
	south
	west
)

type instruction struct {
	left bool
	dist int
}

type pos struct {
	x, y int
}

type state struct {
	pos
	d int
}

func main() {
	instructions := parse("input.txt")

	// part 1
	fmt.Println(p1(instructions))

	// part 2
	fmt.Println(p2(instructions))
}

func p1(instructions []instruction) int {
	s := state{}
	for _, ins := range instructions {
		turn(&s, ins.left)
		walk(&s, ins.dist)
	}
	return abs(s.x) + abs(s.y)
}

func p2(instructions []instruction) int {
	s := state{}
	visited := make(map[pos]bool)
	visited[s.pos] = true
	for _, ins := range instructions {
		turn(&s, ins.left)
		for i := 0; i < ins.dist; i++ {
			walk(&s, 1)
			if visited[s.pos] {
				return abs(s.x) + abs(s.y)
			} else {
				visited[s.pos] = true
			}
		}

	}
	return -1
}

func turn(state *state, left bool) {
	switch left {
	case true:
		state.d = (state.d + 1) % 4
	case false:
		state.d = (state.d + 3) % 4
	}
}

func walk(state *state, dist int) {
	switch state.d {
	case north:
		state.y += dist
	case south:
		state.y -= dist
	case east:
		state.x += dist
	case west:
		state.x -= dist
	}
}

func parse(filename string) []instruction {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	s := strings.Split(strings.TrimSpace(string(b)), ", ")

	instructions := []instruction{}

	for _, ins := range s {
		left := ins[0] == 'L'
		dist, err := strconv.Atoi(ins[1:])
		if err != nil {
			log.Fatal(err)
		}
		instructions = append(instructions, instruction{left, dist})
	}
	return instructions
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
