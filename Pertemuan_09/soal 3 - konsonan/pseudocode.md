program 
kamus
    karakter, tipe : char
	vokal : boolean
algoritma
    input(karakter)
    karakter = lower(karakter)
    tipe <- "Bukan konsonan"
    vokal = karakter == "a" OR karakter == "e" OR karakter == "i" OR karakter == "o" OR karakter == "u"
    if karakter >= "a" && karakter <= "z" && !vokal then
        tipe = "Konsonan"
    endif
    output(tipe)
endprogram