package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		printInt(0)
		z01.PrintRune('\n')
		return
	}

	n := 0
	for _, c := range os.Args[1] {
		if c < '0' || c > '9' {
			printInt(0)
			z01.PrintRune('\n')
			return
		}
		n = n*10 + int(c-'0')
	}

	if n <= 0 {
		printInt(0)
		z01.PrintRune('\n')
		return
	}

	sum := 0
	for i := 2; i <= n; i++ {
		if prime(i) {
			sum += i
		}
	}
	printInt(sum)
	z01.PrintRune('\n')
}

func prime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func printInt(n int) {
	if n > 9 {
		printInt(n / 10)
	}
	z01.PrintRune(rune(n%10 + '0'))
}
