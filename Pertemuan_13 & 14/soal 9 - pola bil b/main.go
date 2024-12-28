package main

import "fmt"

func main() {
	var x, i, j int
	fmt.Scan(&x)

	for i = 1; i <= x; i++ { // Loop untuk baris
		for j = 1; j <= x; j++ { // Loop untuk kolom
			if i == 1 || i == x || j == 1 || j == x {
				fmt.Print(i, " ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}
