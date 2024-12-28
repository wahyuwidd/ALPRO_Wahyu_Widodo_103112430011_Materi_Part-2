program Prima
kamus
    x, i : integer
	prima : boolean
algoritma
    input(x)
    prima <- true
    if x  <= 1 then
        prima <- false
    else
        for i = 2*2 to x do
            if x mod i == 0 do
                prima <- false
                stop
            endif
        endfor
    endif
endprogram