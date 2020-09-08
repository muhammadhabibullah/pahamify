package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Algorithm problems:\n" +
		"1. Print Oscillated Integer Number\n" +
		"2. Bonus coklat\n" +
		"3. Mencari Elemen Terkecil Yang Lebih Besar Dari Setiap Elemen di Sebuah Array\n")

	fmt.Print("Enter problem number you want to solve (1-3): ")
	var problemNum int
	fmt.Scanf("%d", &problemNum)

	switch problemNum {
	case 1:
		fmt.Print("Enter a number between 0 and 1000: ")
		var num int
		fmt.Scanf("%d", &num)
		PrintOscillatedIntNum(num)
	case 2:
		fmt.Print("Masukkan jumlah coklat: ")
		var num int
		fmt.Scanf("%d", &num)
		ChocolateBonus(num)
	case 3:
		fmt.Println("Masukkan set bilangan integer yang dibatasi dengan spasi (contoh: '1 3 2'): ")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		text = strings.TrimSuffix(text, "\n")
		textArray := strings.Split(text, " ")
		intArray := make([]int, len(textArray))
		for i, t := range textArray {
			var err error
			intArray[i], err = strconv.Atoi(t)
			if err != nil {
				log.Fatalf("input bukan bilangan valid: %s", t)
			}
		}
		FindGreatestArrElemAndPrint(intArray)
	}
}
