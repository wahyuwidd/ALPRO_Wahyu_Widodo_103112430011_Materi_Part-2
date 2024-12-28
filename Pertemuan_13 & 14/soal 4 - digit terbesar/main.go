package main

import "fmt"

func main() {
	var bilangan, digitTerbesar, digit int
	fmt.Scan(&bilangan)

	digitTerbesar = 0

	for bilangan > 0 {
		digit = bilangan % 10
		if digit > digitTerbesar {
			digitTerbesar = digit
		}
		bilangan = bilangan / 10
	}

	fmt.Println(digitTerbesar)
}
