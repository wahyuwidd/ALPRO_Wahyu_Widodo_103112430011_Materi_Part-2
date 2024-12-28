program PolaBilanganB
kamus
    x, i, j : int
algoritma
    input(x)
    for i <- 1 to x do
        for j <- 1 to x do
            if i == 1 OR i == x OR j == 1 OR j == x then
				output(i, " ")
			else
				output("  ")
			endif
        endfor
        output(newLine)
    endfor
endprogram