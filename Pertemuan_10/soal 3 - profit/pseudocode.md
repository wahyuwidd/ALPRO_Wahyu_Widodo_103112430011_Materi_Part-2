program Profit
kamus
    keuntunganBulanIni, keuntunganBulanSebelumnya, selisih : real
algoritma
    input(keuntunganBulanIni, keuntunganBulanSebelumnya)
    selisih = keuntunganBulanIni - keuntunganBulanSebelumnya
    if selisih > 0 then
		output("Peningkatan sebesar: ", -selisih)
	else if selisih < 0 then
		output("Penurunan sebesar: ", -selisih)
	else then
		output("Tetap")
	endif
endprogram