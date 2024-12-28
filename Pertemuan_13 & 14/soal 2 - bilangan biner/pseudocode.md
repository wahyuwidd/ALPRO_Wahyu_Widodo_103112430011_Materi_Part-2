program BilanganBiner
kamus
    desimal, sisa : integer
	biner : char
algoritma
    input(desimal)
    biner <- ""
    if desimal == 0 then
        biner = "0"
    else
        while desimal < 0 do
            sisa <- desimal mod 2
            biner <- konversiIntegerKeString(sisa) + biner
            desimal <- desimal div 2
        endwhile
    endif
    output(biner)
endprogram