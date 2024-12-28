package main

import "fmt"

func main() {
	var n, lebar, maxLebar, i int
	fmt.Scan(&n)

	maxLebar = 0

	for i = 0; i < n; i++ {
		fmt.Scan(&lebar)
		if lebar > maxLebar {
			maxLebar = lebar
		}
	}

	fmt.Println(maxLebar)
}
