package main

import "fmt"

func main() {
	var totalBelanja, hargaBayar int
	var bersediaKartu, dapatKartu, dapatDiskon, dapatCashback bool

	fmt.Scan(&totalBelanja, &bersediaKartu)

	dapatKartu = bersediaKartu
	dapatDiskon = totalBelanja >= 100000

	hargaBayar = totalBelanja
	if dapatDiskon {
		hargaBayar = int(float64(totalBelanja) * 0.9)
	}

	dapatCashback = hargaBayar >= 200000 && dapatKartu

	if dapatCashback {
		hargaBayar -= 75000
	}

	fmt.Println("Kartu?", dapatKartu)
	fmt.Println("Diskon?", dapatDiskon)
	fmt.Println("Cashback?", dapatCashback)
	fmt.Println("Rp.", hargaBayar)
}
