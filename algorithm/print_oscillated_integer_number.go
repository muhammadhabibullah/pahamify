package main

import (
	"fmt"
)

func PrintOscillatedIntNum(n int) {
	if n < 1000 && n > 0 {
		printN(n, n, false)
		fmt.Println()
	} else {
		fmt.Println("Input number must be a number between 0 and 1000")
	}
}

func printN(n, origin int, redo bool) {
	fmt.Printf("%d ", n)

	if !redo {
		n -= 5
	} else {
		n += 5
	}

	if n <= origin {
		if n > 0 {
			printN(n, origin, redo)
		} else {
			printN(n, origin, !redo)
		}
	}
}
