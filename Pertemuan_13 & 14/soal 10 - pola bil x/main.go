package main

import "fmt"

func main() {
	var x, i, j int
	fmt.Scan(&x)

	for i = 1; i <= x; i++ {
		for j = 1; j <= x; j++ {
			if j == i || j == x-i+1 { // Kondisi untuk membuat "X"
				fmt.Print(i) // Cetak angka sesuai baris
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println() // Pindah ke baris berikutnya
	}
}
