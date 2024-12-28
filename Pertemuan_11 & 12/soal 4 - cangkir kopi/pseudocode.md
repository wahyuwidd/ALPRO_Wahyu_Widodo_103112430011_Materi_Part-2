program CangkirKopi
kamus
    n, m, x, y, jumlahCangkir : int
algoritma
    input(n, m, x, y)
    jumlahCangkir <- 0
    while n >= x AND m >= y do
        n -= x
		m -= y
		jumlahCangkir <- jumlahCangkir + 1 
    endwhile
endprogram