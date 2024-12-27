package main

import "fmt"

func main() {
	var jumlahGagal int
	var username, password string
	jumlahGagal = 0

	fmt.Scan(&username, &password)
	for username != "admin" || password != "admin" {

		fmt.Scan(&username, &password)

		jumlahGagal++
	}

	fmt.Println(jumlahGagal)
	fmt.Println("Login berhasil")
}
