package main

import "fmt"

func main() {
	var T, V, E int
	var selesai bool
	V = 0
	fmt.Scan(&T)
	for selesai = false; !selesai; {
		fmt.Scan(&E)
		V = V + E
		if V >= T {
			fmt.Println(true)
		} else {
			fmt.Println(false)
		}
		selesai = V >= T
	}
}
