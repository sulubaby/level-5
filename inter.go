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
	done := ""
	for _, c := range s1 {
		if contains(s2, c) && !contains(done, c) {
			z01.PrintRune(c)
			done += string(c)
		}
	}
	z01.PrintRune('\n')
}

func contains(s string, r rune) bool {
	for _, c := range s {
		if c == r {
			return true
		}
	}
	return false
}
