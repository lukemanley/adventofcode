package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)

	scanner.Split(bufio.ScanLines)

	count := 0
	count2 := 0

	for scanner.Scan() {

		line := scanner.Text()

		var r2 rune
		var r3 rune

		repeat := false
		for i, r := range line {

			if i > 0 && !repeat && r == r2 {
				repeat = true
			}
			r2 = r

		}
		if repeat {

			if !strings.Contains(line, "ab") && !strings.Contains(line, "cd") && !strings.Contains(line, "pq") && !strings.Contains(line, "xy") {

				vowels := 0
				vowels += strings.Count(line, "a")
				vowels += strings.Count(line, "e")
				vowels += strings.Count(line, "i")
				vowels += strings.Count(line, "o")
				vowels += strings.Count(line, "u")
				if vowels > 2 {
					count++
				}
			}
		}

		pair := false
		repeat = false

		for i, r := range line {
			if i > 0 && !pair {
				s := strings.Split(line, fmt.Sprint(string(r2), string(r)))
				if len(s) > 2 {
					pair = true
				}
			}
			if i > 1 && !repeat {
				if r == r3 {
					repeat = true
				}
			}
			r3 = r2
			r2 = r
		}
		if pair && repeat {
			count2++
		}
	}
	fmt.Println("nice string count1:", count)
	fmt.Println("nice string count2:", count2)

}
