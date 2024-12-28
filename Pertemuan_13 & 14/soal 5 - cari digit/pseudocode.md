program CariDigit
kamus
    x, n, digit : integer
	ada : boolean
algoritma
    input(x, n)
    ada <- false
    while n > 0 do
        digit <- n mod 10
        if digit == x then
            ada <- true
            stop
        endif
        n = n div 10
    endwhile
    output(ada)
endprogram