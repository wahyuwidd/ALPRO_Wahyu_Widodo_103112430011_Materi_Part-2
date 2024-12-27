package main

import (
	"fmt"
)

func main() {
	var keuntunganBulanIni, keuntunganBulanSebelumnya, selisih float64

	fmt.Scan(&keuntunganBulanIni, &keuntunganBulanSebelumnya)

	selisih = keuntunganBulanIni - keuntunganBulanSebelumnya

	if selisih > 0 {
		fmt.Printf("Peningkatan sebesar %.0f\n", -selisih) // Format output agar tanpa desimal
	} else if selisih < 0 {
		fmt.Printf("Penurunan sebesar %.0f\n", -selisih) // math.Abs untuk nilai absolut
	} else {
		fmt.Println("Tetap")
	}
}
