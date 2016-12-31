package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

func main() {

	input := []byte("iwrupvqb")

	pre := "00000"
	i := 1

	for {
		b := append(input, []byte(strconv.Itoa(i))...)
		h := md5.Sum(b)
		s2 := fmt.Sprintf("%x", h)

		if strings.HasPrefix(s2, pre) {
			switch pre {
			case "00000":
				fmt.Println("Solution 1:", i)
				pre = "000000"
				continue
			case "000000":
				fmt.Println("Solution 2:", i)
				return
			}
		}
		i++
	}
}
