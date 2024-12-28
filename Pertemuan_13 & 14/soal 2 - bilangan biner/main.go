package main

import (
	"fmt"
	"strconv"
)

func main() {
	var desimal, sisa int
	var biner string
	fmt.Scan(&desimal)

	biner = ""

	if desimal == 0 {
		biner = "0"
	} else {
		for desimal > 0 {
			sisa = desimal % 2
			biner = strconv.Itoa(sisa) + biner
			desimal /= 2
		}
	}

	fmt.Println(biner)
}
