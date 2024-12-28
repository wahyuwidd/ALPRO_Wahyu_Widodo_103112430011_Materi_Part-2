program TangkiAir
kamus
    T, V, E : integer
	selesai : boolean
algoritma
    V <- 0
    input(T)
    repeat 
        input(E)
        V <- V + E
        if V >= T then
            output(false)
        else 
            output(true)
        endif
    until V >= T
endprogram