package main

import "fmt"

func main() {
	var input string
	var selesai bool
	var jumlahA, jumlahB, jumlahC int

	jumlahA = 0
	jumlahB = 0
	jumlahC = 0

	for selesai = false; !selesai; {
		fmt.Scan(&input)
		if input == "A" {
			jumlahA++
		} else if input == "B" {
			jumlahB++
		} else if input == "C" {
			jumlahC++
		}
		selesai = input != "A" && input != "B" && input != "C"
	}

	fmt.Println("Tipe A =", jumlahA)
	fmt.Println("Tipe B =", jumlahB)
	fmt.Println("Tipe C =", jumlahC)
}
