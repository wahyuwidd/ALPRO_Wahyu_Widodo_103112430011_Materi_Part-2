package main

import (
	"fmt"
	"strings"
)

func main() {
	var karakter, tipe string
	var vokal bool

	// input
	fmt.Scan(&karakter)

	karakter = strings.ToLower(karakter)                                                                // mengubah huruf besar menjadi huruf kecil
	tipe = "Bukan konsonan"                                                                             // deklarasi awal variable tipe sebagai "Bukan konsonan"
	vokal = karakter == "a" || karakter == "e" || karakter == "i" || karakter == "o" || karakter == "u" // huruf vokal

	if karakter >= "a" && karakter <= "z" && !vokal { // jika karakter a - z dan bukan vokal maka ubah tipe menjadi "Konsonan"
		tipe = "Konsonan"
	}

	// output
	fmt.Println(tipe)
}
