program Konsekutif
kamus
    n, digitTerakhir, digitSekarang : integer
	konsekutif boolean
algoritma
    input(n)
    digitTerakhir <- n mod 10
	n <- n / 10
	konsekutif <- true
    while n > 0 do
		digitSekarang <- n mod 10

		if digitSekarang not digitTerakhir-1 AND digitSekarang not digitTerakhir+1 then
			konsekutif <- false
		endif

		digitTerakhir <- digitSekarang
		n <- n / 10
	endwhile

endprogram