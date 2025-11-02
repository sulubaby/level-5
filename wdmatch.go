package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	s1, s2 := os.Args[1], os.Args[2]
	i := 0
	for _, c := range s2 {
		if i < len(s1) && rune(s1[i]) == c {
			i++
		}
	}
	if i == len(s1) {
		for _, c := range s1 {
			z01.PrintRune(c)
		}
		z01.PrintRune('\n')
	}
}
