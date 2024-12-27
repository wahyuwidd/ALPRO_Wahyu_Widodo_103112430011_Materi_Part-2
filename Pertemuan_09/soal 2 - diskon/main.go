package main

import "fmt"

func main() {
	var uang_awal, uang_akhir, diskon float64
	var ikut bool

	// input
	fmt.Scan(&uang_awal, &ikut)

	// proses jika ikut assesment maka mendapatkan diskon
	if ikut {
		diskon = (35 / 100) * uang_awal
	}

	uang_akhir = uang_awal - diskon

	// output
	fmt.Print(uang_akhir)
}
