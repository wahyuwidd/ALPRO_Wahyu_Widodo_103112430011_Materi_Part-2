package main

import "fmt"

func main() {
	var bilangan int

	// input
	fmt.Scan(&bilangan)

	// jika bilangan habis dibagi 3 maka kelipatan 3
	if bilangan%3 == 0 {
		fmt.Println("Kelipatan 3")
	}

	// jika bilangan habis dibagi 5 maka kelipatan 5
	if bilangan%5 == 0 {
		fmt.Print("Kelipatan 5")
	}
}
