program 
kamus
    sisi1, sisi2, sisi3 : integer
algoritma
    input(sisi1, sisi2, sisi3)
    if sisi1 == sisi2 AND sisi2 == sisi3  then
		output("Segitiga Sama Sisi")
	else if sisi1 == sisi2 OR sisi1 == sisi3 OR sisi2 == sisi3 then
		output("Segitiga Sama Kaki")
	else then
		output("Segitiga Sembarang")
	endif
endprogram