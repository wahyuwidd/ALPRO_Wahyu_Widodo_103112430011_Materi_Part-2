package main

import "fmt"

func main() {
	var x, i, j int
	fmt.Scan(&x)

	for i = 1; i <= x; i++ {
		for j = 1; j <= 2*x-1; j++ {
			if j == x-i+1 || j == x+i-1 || (i == x/2+1 && j > x-i+1 && j < x+i-1 && x%2 != 0) {
				fmt.Print(i)
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
