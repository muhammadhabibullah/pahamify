package main

import (
	"fmt"
)

func ChocolateBonus(n int) {
	fmt.Println(countChocolateBonus(n))
}

func countChocolateBonus(n int) int {
	if n <= 0 {
		return 0
	}
	bonus := n / 5
	return n + countChocolateBonus(bonus)
}
