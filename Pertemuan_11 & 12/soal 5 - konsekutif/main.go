package main

import "fmt"

func main() {
	var n, digitTerakhir, digitSekarang int
	var konsekutif bool

	fmt.Scan(&n)

	digitTerakhir = n % 10
	n = n / 10

	konsekutif = true

	for n > 0 {
		digitSekarang = n % 10

		if digitSekarang != digitTerakhir-1 && digitSekarang != digitTerakhir+1 {
			konsekutif = false
		}

		digitTerakhir = digitSekarang
		n = n / 10
	}
	fmt.Println(konsekutif)

}
