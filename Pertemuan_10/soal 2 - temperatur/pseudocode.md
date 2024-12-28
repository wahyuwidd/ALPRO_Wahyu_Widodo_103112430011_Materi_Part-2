program Temperatur
kamus
    temp1, temp2, temp3, temp4, temp5 : real
    stabil, naik, turun : bool
algoritma
    input(temp1, temp2, temp3, temp4, temp5)
    stabil <- true
	naik <- false
	turun <- false
    if temp2 > temp1 then
		naik <- true
	else if temp2 < temp1
		turun <- true
	endif
    if naik then
		if temp3 < temp2 OR temp4 < temp3 OR temp5 < temp4 then
			stabil <- false
		endif
	else if turun
		if temp3 > temp2 OR temp4 > temp3 OR temp5 > temp4 then
			stabil <- false
		endif
	endif
    if stabil then
		output("Stabil naik/turun")
	else 
		output("Tidak stabil")
	endif
endprogram