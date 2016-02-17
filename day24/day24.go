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
	var weights []int

	for s.Scan() {
		w, err := strconv.Atoi(s.Text())
		if err != nil {
			log.Fatal(err)
		}
		weights = append(weights, w)
	}

	total := 0
	for _, w := range weights {
		total += w
	}

	reserved := make(map[int]bool)
	_, minprod1 := solve(weights, reserved, 0, 0, 0, 1, total/3, len(weights)/3)
	_, minprod2 := solve(weights, reserved, 0, 0, 0, 1, total/4, len(weights)/4)

	fmt.Println("Solution 1:", minprod1)
	fmt.Println("Solution 2:", minprod2)
}

func solve(weights []int, reserved map[int]bool, start, sum, count, prod, target, maxcount int) (int, int) {

	if sum > target {
		return -1, -1
	}
	if sum == target {
		return count, prod
	}

	mincount := -1
	minprod := -1

	for i, w := range weights[start:] {
		if len(reserved) == maxcount {
			continue
		}
		reserved[start+i] = true

		mincount2, minprod2 := solve(weights, reserved, start+i+1, sum+w, count+1, prod*w, target, maxcount)

		balanced := false
		if mincount2 > 0 {
			balanced = balances(weights, reserved, 0, target)
		}

		delete(reserved, start+i)

		if !balanced {
			continue
		}

		if mincount == -1 || mincount2 < mincount {
			mincount = mincount2
			minprod = minprod2
			continue
		}

		if mincount2 == mincount {
			if minprod2 < minprod {
				minprod = minprod2
			}
		}

	}
	return mincount, minprod
}

func balances(weights []int, reserved map[int]bool, sum int, target int) bool {
	if sum == target {
		return true
	}
	if sum > target {
		return false
	}

	for i, w := range weights {
		used, _ := reserved[i]
		if used {
			continue
		}

		reserved[i] = true
		b := balances(weights, reserved, sum+w, target)
		delete(reserved, i)

		if b {
			return true
		}
	}
	return false
}
