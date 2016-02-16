package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type transform struct {
	old string
	new string
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)

	molecule := ""

	var transforms []*transform

	for s.Scan() {
		line := s.Text()
		if line == "" {
			s.Scan()
			molecule = s.Text()
		}

		tf := transform{}

		fmt.Sscanf(line, "%s => %s", &tf.old, &tf.new)
		transforms = append(transforms, &tf)
	}

	ms := make(map[string]bool)

	fmt.Println("Solution 1:", p1(molecule, transforms, ms))
	fmt.Println("Solution 2:", step(molecule, transforms, 0))
}

func p1(s string, transforms []*transform, ms map[string]bool) int {

	for _, tf := range transforms {
		sli := strings.Split(s, tf.old)
		for i, _ := range sli {
			if i == len(sli)-1 {
				continue
			}
			var new []string
			for i2, str := range sli {
				new = append(new, str)
				if i2 == len(sli)-1 {
					continue
				}
				if i == i2 {
					new = append(new, tf.new)
				} else {
					new = append(new, tf.old)
				}
			}
			newstr := strings.Join(new, "")
			ms[newstr] = true
		}
	}

	return len(ms)
}

func step(s string, transforms []*transform, n int) int {
	if s == "e" {
		return n
	}
	for _, tf := range transforms {
		for i := strings.Index(s, tf.new); i != -1; {
			s2 := s[:i] + strings.Replace(s[i:], tf.new, tf.old, 1)
			n2 := step(s2, transforms, n+1)
			if n2 != -1 {
				return n2
			}
		}
	}
	return -1
}
