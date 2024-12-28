program PolaBilanganA
kamus
    x, i, j : integer
algoritma
    input(x)
    for i <- 1 to x do
        for j <- 1 to x do
            if j == x-i+1 OR j == x+i-1 OR (i == x/2+1 AND j > x-i+1 AND j < x+i-1 AND x%2 not 0) then
				output(i)
			else
				output(" ")
			endif
        endfor
        output(newLine)
    endfor

endprogram