package main

import "fmt"

func main() {
	var sisi1, sisi2, sisi3 int

	// input
	fmt.Scan(&sisi1, &sisi2, &sisi3)

	if sisi1 == sisi2 && sisi2 == sisi3 {
		//jika semua sisi sama
		fmt.Println("Segitiga Sama Sisi")
	} else if sisi1 == sisi2 || sisi1 == sisi3 || sisi2 == sisi3 {
		//jika dua sisi sama
		fmt.Println("Segitiga Sama Kaki")
	} else {
		//jika semua sisi berbeda
		fmt.Println("Segitiga Sembarang")
	}
}
