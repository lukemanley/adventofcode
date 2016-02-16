package main

import "fmt"

func main() {
	pw := []byte("cqjxjnds")

	pw = next(pw)
	fmt.Println("Solution 1:", string(pw))

	pw = next(pw)
	fmt.Println("Solution 2:", string(pw))
}

func next(pw []byte) []byte {
	for {
		for i := len(pw) - 1; i > -1; i-- {
			b := pw[i]
			if b == 'z' {
				pw[i] = 'a'
				continue
			}
			pw[i] += 1
			break
		}

		// two sets of non-overlapping pairs
		if !pairs(pw) {
			continue
		}

		// three straight of increasing letters
		if !ascend3(pw) {
			continue
		}

		return pw
	}
}

func pairs(pw []byte) bool {
	n := 0
	i := 1
	for i < len(pw) {
		if pw[i] == pw[i-1] {
			n++
			if n == 2 {
				return true
			}
			i += 2
			continue
		}
		i += 1
	}
	return false
}

func ascend3(pw []byte) bool {
	i := 2
	for i < len(pw) {
		b := pw[i]
		if b == pw[i-1]+1 && b == pw[i-2]+2 {
			return true
		}
		i++
	}
	return false
}
