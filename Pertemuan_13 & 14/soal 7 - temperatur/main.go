package main

import "fmt"

func main() {
	var input, min, max, nolBerturutTurut, total, jumlahInput int
	var selesai bool

	for selesai = false; !selesai; {
		fmt.Scan(&input)
		if input == 0 {
			nolBerturutTurut++
		}
		if input < min {
			min = input
		}
		if input > max {
			max = input
		}

		total = total + input
		jumlahInput++
		selesai = nolBerturutTurut == 2
	}

	fmt.Println(max)
	fmt.Println(min)
	fmt.Println(float64(float64(total) / float64(jumlahInput-1)))
}
