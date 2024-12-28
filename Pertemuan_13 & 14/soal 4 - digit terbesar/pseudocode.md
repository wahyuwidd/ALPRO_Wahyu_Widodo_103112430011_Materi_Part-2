program DigitTerbesar
kamus
    bilangan, digitTerbesar, digit : integer
algoritma
    input(bilangan)
    digitTerbesar <- 0 
    while bilangan > 0 do
        digit <- bilangan mod 10
        if digit > digitTerbesar then
            digitTerbesar <- digit
        endif
        bilangan <- bilangan div 10
    endwhile
    output(digitTerbesar)
endprogram