program Temperatur
kamus
    input, min, max, nolBerturutTurut, total, jumlahInput : integer
	selesai : boolean
algoritma
    repeat
        input(input)
        if input == 0 then
			nolBerturutTurut <- nolBerturutTurut + 1
		endif
		if input < min then
			min <- input
		endif
		if input > max then
			max <- input
		endif
        total <- total + input
		jumlahInput <- jumlahInput + 1
    until nolBerturutTurut == 2
    output(min)
    output(max)
    output(total/jumlahInput-1)
endprogram