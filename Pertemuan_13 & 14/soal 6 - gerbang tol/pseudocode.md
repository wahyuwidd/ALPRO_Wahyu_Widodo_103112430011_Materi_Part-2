program GerbangTol
kamus
    input : char
	selesai : boolean
	jumlahA, jumlahB, jumlahC : integer
algoritma
    jumlahA <- 0
	jumlahB <- 0
	jumlahC <- 0
    repeat
        input(input)
        if input == "A" then
			jumlahA <- jumlahA + 1
		else if input == "B" 
			jumlahB <- jumlahB + 1
		else if input == "C" 
			jumlahC <- jumlahC + 1
		endif
    until input not "A" AND input not "B" AND input not "C"
endprogram