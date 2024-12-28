package main

import "fmt"

func main() {
	var x, totalJumlah, digit int

	fmt.Scan(&x)

	totalJumlah = 0

	for x > 0 {
		digit = x % 10
		fmt.Print(digit, " ")
		totalJumlah += digit
		x = x / 10
	}

	fmt.Print("\n", totalJumlah)
}
