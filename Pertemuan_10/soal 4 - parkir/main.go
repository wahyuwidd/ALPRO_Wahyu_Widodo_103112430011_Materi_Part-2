package main

import "fmt"

func main() {
	var h1, m1, h2, m2 int
	var jam, menit int
	var totalMenitMasuk, totalMenitKeluar, selisihMenit int

	fmt.Scan(&h1, &m1, &h2, &m2)

	totalMenitMasuk = h1*60 + m1
	totalMenitKeluar = h2*60 + m2

	if h2 < h1 {
		totalMenitKeluar += 720
	}

	selisihMenit = totalMenitKeluar - totalMenitMasuk
	jam = selisihMenit / 60
	menit = selisihMenit % 60

	fmt.Printf("%d jam %d menit\n", jam, menit)
}
