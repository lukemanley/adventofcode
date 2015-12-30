package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := []byte("iwrupvqb")

	fiveFound := false

	for i := 1; ; i++ {
		s := strconv.Itoa(i)
		b := append(input, []byte(s)...)
		h := md5.Sum(b)
		s2 := fmt.Sprintf("%x", h)

		if strings.HasPrefix(s2, "00000") {
			if !fiveFound {
				fmt.Println("5 zeros:", i)
				fiveFound = true
			}
			if strings.HasPrefix(s2, "000000") {
				fmt.Println("6 zeros:", i)
				return
			}
		}
	}
}
