package main

import "fmt"

func main() {
	var gol1, gol2, gol3, gol4 int
	var maks1, min1, maks2, min2, maks, min int

	// input
	fmt.Scan(&gol1, &gol2, &gol3, &gol4)

	if gol1 > gol2 {
		maks1 = gol1
		min1 = gol2
	} else {
		maks1 = gol2
		min1 = gol1
	}

	if maks1 > gol3 {
		maks2 = maks1
	} else {
		maks2 = gol3
	}

	if min1 < gol3 {
		min2 = min1
	} else {
		min2 = gol3
	}

	if maks2 > gol4 {
		maks = maks2
	} else {
		maks = gol4
	}

	if min2 < gol4 {
		min = min2
	} else {
		min = gol4
	}

	// output
	fmt.Println(maks, min)
}
