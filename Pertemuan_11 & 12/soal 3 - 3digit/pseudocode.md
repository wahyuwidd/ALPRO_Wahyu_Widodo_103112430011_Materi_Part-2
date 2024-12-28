program 3Digit
kamus
    x, totalJumlah, digit : integer
algoritma
    input(x)
    totalJumlah <- 0
    while x > 0 do
        digit <- x mod 10
        output(digit)
        totalJumlah <- totalJumlah + digit
        x = x div 10
    endwhile
    output(totalJumlah)
endprogram