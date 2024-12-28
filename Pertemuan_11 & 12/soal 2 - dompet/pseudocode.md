program Dompet
kamus
    transaksi, saldo : integer
algoritma
    saldo <- 0
    input(transaksi)
    while (true) do
        input(transaksi)
        if transaksi = 0 then
            stop
        endif
        saldo <- saldo + transaksi
    endwhile
    output(saldo)
endprogram