program PolaBilanganX
kamus
    x, i, j : integer
algoritma
    input(x)
    for i <- 1 to x do
        for j <- 1 to x do
            if j == i OR j == x-i+1 then
                output(i)
            else
                output("  ")
            endif
        endfor
        output(newLine)
    endfor 
endprogram