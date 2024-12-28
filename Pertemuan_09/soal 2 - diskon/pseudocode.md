program Diskon
kamus
    uang_awal, uang_akhir, diskon : real
	ikut : boolean
algoritma
    input(uang_awal, ikut)
    if ikut == true then
        diskon <- (35 / 100) * uang_awal
    endif
    uang_akhir <- uang_awal - diskon
    output(uang_akhir)
endprogram