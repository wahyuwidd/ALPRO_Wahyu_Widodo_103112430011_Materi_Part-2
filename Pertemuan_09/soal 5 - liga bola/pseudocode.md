program LigaBola
kamus
    gol1, gol2, gol3, gol4 integer
	maks1, min1, maks2, min2, maks, min integer
algoritma
    input(gol1, gol2, gol3, gol4)
    if gol1 > gol2 then
        maks1 <- gol1
        min1 <- gol2
    else
        maks1 <- gol2
        min1 <- gol1
    endif

    if maks1 > gol3 then
        maks2 <- maks1
    else
        maks2 <- gol3
    endif

    if min1 < gol3 then
        min2 <- min1
    else
        min2 <- gol3
    endif

    if maks2 > gol4 then
        maks <- maks2
    else
        maks <- gol4
    endif

    if min2 < gol4 then
        min <- min2
    else
        min <- gol4
    endif

    output(min,max)
endprogram