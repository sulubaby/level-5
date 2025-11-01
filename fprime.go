package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n < 2 {
		return
	}
	i := 2
	for i*i <= n {
		if n%i == 0 {
			fmt.Print(i)
			n /= i
			if n > 1 {
				fmt.Print("*")
			}
		} else {
			i++
		}
	}
	if n > 1 {
		fmt.Print(n)
	}
	fmt.Println()
}
