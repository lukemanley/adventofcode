package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)

	var instructions []string

	for s.Scan() {
		instructions = append(instructions, s.Text())
	}

	fmt.Println("solution 1:", process(instructions, make(map[byte]int)))
	fmt.Println("solution 2:", process(instructions, map[byte]int{'a': 1}))
}

func process(instructions []string, regs map[byte]int) int {
	for i := 0; i < len(instructions); i++ {
		ins := instructions[i]

		switch ins[:3] {
		case "hlf":
			regs[ins[4]] /= 2

		case "tpl":
			regs[ins[4]] *= 3

		case "inc":
			regs[ins[4]]++

		case "jmp":
			j, err := strconv.Atoi(ins[4:])
			if err != nil {
				log.Fatal("invalid instruction: " + ins)
			}
			i += j - 1

		case "jie":
			if regs[ins[4]]%2 == 1 {
				continue
			}
			j, err := strconv.Atoi(ins[7:])
			if err != nil {
				log.Fatal("invalid instruction: " + ins)
			}
			i += j - 1

		case "jio":
			if regs[ins[4]] != 1 {
				continue
			}
			j, err := strconv.Atoi(ins[7:])
			if err != nil {
				log.Fatal("invalid instruction: " + ins)
			}
			i += j - 1

		default:
			log.Fatal("invalid instruction: " + ins)
		}
	}
	return regs['b']
}
