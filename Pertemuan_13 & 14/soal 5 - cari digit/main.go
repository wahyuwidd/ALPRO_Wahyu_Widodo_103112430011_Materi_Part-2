package main

import "fmt"

func main() {
	var x, n, digit int
	var ada bool
	fmt.Scan(&x, &n)

	ada = false

	for n > 0 {
		digit = n % 10
		if digit == x {
			ada = true
			break
		}
		n = n / 10
	}

	fmt.Println(ada)
}
