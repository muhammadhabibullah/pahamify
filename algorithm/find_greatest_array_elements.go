package main

import (
	"fmt"
)

func FindGreatestArrElemAndPrint(arr []int) {
	var max, maxIdx int
	for idx, a := range arr {
		if a > max {
			max = a
			maxIdx = idx
		}
	}

	for idx, a := range arr {
		if idx == maxIdx {
			fmt.Print("^ ")
		} else {
			printLargerInt(a)
		}
	}

	fmt.Println()
}

func printLargerInt(n int) {
	fmt.Printf("%d ", n+1)
}
