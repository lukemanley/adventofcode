package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	var instrux []string
	for s.Scan() {
		instrux = append(instrux, s.Text())
	}

	c := newCircuit()
	fmt.Printf("Solution 1: %d\n", c.run(instrux))

	c2 := newCircuit()
	c2.wires["b"] = 3176
	fmt.Printf("Solution 2: %d\n", c2.run(instrux))
}

type circuit struct {
	wires map[string]uint16
}

func newCircuit() *circuit {
	return &circuit{wires: make(map[string]uint16)}
}

func (c *circuit) run(instrux []string) uint16 {
	for {
		for _, s := range instrux {
			p := strings.Split(s, " ")
			dest := p[len(p)-1]
			if _, ok := c.wires[dest]; ok {
				continue
			}
			var i uint16
			var err error

			switch {
			case len(p) == 3:
				i, err = getOrParse(p[0], c)
				if err != nil {
					continue
				}
			case p[0] == "NOT":
				i, err = getOrParse(p[1], c)
				if err != nil {
					continue
				}
				i = ^i
			case p[1] == "AND":
				i1, err := getOrParse(p[0], c)
				if err != nil {
					continue
				}
				i2, err := getOrParse(p[2], c)
				if err != nil {
					continue
				}
				i = i1 & i2
			case p[1] == "OR":
				i1, err := getOrParse(p[0], c)
				if err != nil {
					continue
				}
				i2, err := getOrParse(p[2], c)
				if err != nil {
					continue
				}
				i = i1 | i2
			case p[1] == "RSHIFT":
				i1, err := getOrParse(p[0], c)
				if err != nil {
					continue
				}
				i2, err := getOrParse(p[2], c)
				if err != nil {
					continue
				}
				i = i1 >> i2
			case p[1] == "LSHIFT":
				i1, err := getOrParse(p[0], c)
				if err != nil {
					continue
				}
				i2, err := getOrParse(p[2], c)
				if err != nil {
					continue
				}
				i = i1 << i2
			default:
				panic("instruction not recognized")
			}
			c.wires[dest] = i
		}
		a, ok := c.wires["a"]
		if ok {
			return a
		}
	}
}

func getOrParse(s string, c *circuit) (uint16, error) {
	i, ok := c.wires[s]
	if ok {
		return i, nil
	}
	i64, err := strconv.ParseUint(s, 10, 16)
	return uint16(i64), err
}
