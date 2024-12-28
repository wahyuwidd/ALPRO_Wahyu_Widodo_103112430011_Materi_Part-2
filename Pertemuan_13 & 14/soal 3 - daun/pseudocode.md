program Daun
kamus
    n, lebar, maxLebar, i : integer
algoritma
    input(n)
    maxLebar <- 0
    for i <- 1 to n do
        input(lebar)
        if lebar > maxLebar then
            maxLebar <- lebar
        endif
    endfor
    output(maxLebar)
endprogram