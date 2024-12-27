package main

import "fmt"

func main() {
	var n, kapasitas_mobil, jumlah_mobil int

	// nilai awal 7 yang merupakan kapasitas mobil
	kapasitas_mobil = 7

	// input
	fmt.Scan(&n)

	// proses
	jumlah_mobil = n / kapasitas_mobil
	if n%kapasitas_mobil != 0 {
		jumlah_mobil = jumlah_mobil + 1
	}

	// output
	fmt.Print(jumlah_mobil)
}
