program Tamasya
kamus
    n, kapasitas_mobil, jumlah_mobil : integer
algoritma
    kapasitas_mobil <- 7
    input(n)
    jumlah_mobil <- n div kapasitas_mobil
    if n div kapasitas_mobil != 0 then
        jumlah_mobil <- jumlah_mobil + 1
    endif
    output(jumlah_mobil)
endprogram