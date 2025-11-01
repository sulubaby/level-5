package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	for _, arg := range os.Args[1:] {
		for i := 0; i < len(arg); i++ {
			c := arg[i]
			next := byte(' ')
			if i+1 < len(arg) {
				next = arg[i+1]
			}
			if c >= 'A' && c <= 'Z' {
				c += 32
			}
			if c >= 'a' && c <= 'z' && (next == ' ' || next == '\t' || next == '\n') {
				c -= 32
			}
			z01.PrintRune(rune(c))
		}
		z01.PrintRune('\n')
	}
}
