package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var re = regexp.MustCompile(`(\w+) would (\w+) (\d+) happiness units by sitting next to (\w+).`)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	m := make(map[string]int)
	peopleset := make(map[string]bool)

	s := bufio.NewScanner(f)

	for s.Scan() {
		match := re.FindStringSubmatch(s.Text())
		peopleset[match[1]] = true
		k := match[1] + "_" + match[4]
		v, err := strconv.Atoi(match[3])
		if err != nil {
			log.Fatal(err)
		}
		if match[2] == "lose" {
			v = -v
		}
		m[k] = v
	}

	options := make(map[string]bool)

	toAssign := ""
	for name, _ := range peopleset {
		if toAssign != "" {
			toAssign += "_"
		}
		toAssign += name
	}

	assignSeats(options, "", toAssign)

	max := 0
	max2 := 0

	for opt, _ := range options {
		optTotal := 0
		minLink := 0
		sli := strings.Split(opt, "_")
		for i, name := range sli {
			left := ""
			right := ""
			if i == 0 {
				left = sli[len(sli)-1]
			} else {
				left = sli[i-1]
			}
			if i == len(sli)-1 {
				right = sli[0]
			} else {
				right = sli[i+1]
			}
			h1, ok := m[name+"_"+left]
			if !ok {
				panic("utility not found: " + name + " " + left)
			}
			h2, ok := m[name+"_"+right]
			if !ok {
				panic("utility not found: " + name + " " + right)
			}
			optTotal += h1 + h2

			h1b, ok := m[left+"_"+name]
			if !ok {
				panic("utility no found")
			}
			h2b, ok := m[right+"_"+name]
			if !ok {
				panic("utility not found")
			}

			if i == 0 || h1+h1b < minLink {
				minLink = h1 + h1b
			}

			if h2+h2b < minLink {
				minLink = h2 + h2b
			}
		}
		if optTotal > max {
			max = optTotal
		}
		if optTotal-minLink > max2 {
			max2 = optTotal - minLink
		}
	}
	fmt.Println(max)
	fmt.Println(max2)
}

func assignSeats(options map[string]bool, assigned, toAssign string) {
	if toAssign == "" {
		options[assigned] = true
		return
	}
	for _, name := range strings.Split(toAssign, "_") {
		assignSeats(options, setAdd(assigned, name), setRemove(toAssign, name))
	}
	return
}

func setAdd(s, s2 string) string {
	if s == "" {
		return s2
	}
	return s + "_" + s2
}

func setRemove(s, r string) string {
	if s == r {
		return ""
	}
	if strings.HasPrefix(s, r) {
		return strings.Replace(s, r+"_", "", 1)
	}
	return strings.Replace(s, "_"+r, "", 1)
}
