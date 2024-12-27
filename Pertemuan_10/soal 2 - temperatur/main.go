package main

import "fmt"

func main() {
	var temp1, temp2, temp3, temp4, temp5 float64
	var stabil, naik, turun bool

	fmt.Scan(&temp1, &temp2, &temp3, &temp4, &temp5)

	stabil = true
	naik = false
	turun = false

	if temp2 > temp1 {
		naik = true
	} else if temp2 < temp1 {
		turun = true
	}

	if naik {
		if temp3 < temp2 || temp4 < temp3 || temp5 < temp4 {
			stabil = false
		}
	} else if turun {
		if temp3 > temp2 || temp4 > temp3 || temp5 > temp4 {
			stabil = false
		}
	}

	if stabil {
		fmt.Println("Stabil naik/turun")
	} else {
		fmt.Println("Tidak stabil")
	}
}
