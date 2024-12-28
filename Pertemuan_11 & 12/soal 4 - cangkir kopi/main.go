package main

import "fmt"

func main() {
	var n, m, x, y, jumlahCangkir int

	fmt.Scan(&n, &m, &x, &y)

	jumlahCangkir = 0

	for n >= x && m >= y {
		n -= x
		m -= y
		jumlahCangkir++
	}

	fmt.Println(jumlahCangkir)
}
