program AkhirTahun
kamus
    totalBelanja, hargaBayar : integer
    bersediaKartu, dapatKartu, dapatDiskon, dapatCashback : boolean
algoritma
    input(&totalBelanja, &bersediaKartu)
    dapatKartu <- bersediaKartu
	dapatDiskon <- totalBelanja >= 100000
    hargaBayar <- totalBelanja
    if dapatDiskon then
		hargaBayar <- int(float64(totalBelanja) * 0.9)
	else
        harga_bayar <- total_belanja
    endif
    dapat_cashback <- (harga_bayar >= 200000 AND dapat_kartu)
    if dapatCashback then
		hargaBayar -= 75000
	endif
    output("Kartu?", dapatKartu)
	output("Diskon?", dapatDiskon)
	output("Cashback?", dapatCashback)
	output("Rp.", hargaBayar)
endprogram