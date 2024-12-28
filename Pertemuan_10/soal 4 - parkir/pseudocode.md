program 
kamus
    h1, m1, h2, m2 integer
	jam, menit integer
    totalMenitMasuk, totalMenitKeluar, selisihMenit integer
algoritma
    input(h1, m1, h2, m2)
    totalMenitMasuk <- h1*60 + m1
	totalMenitKeluar <- h2*60 + m2
    if h2 < h1 then
		totalMenitKeluar <- totalMenitKeluar + 720
	endif
    selisihMenit <- totalMenitKeluar - totalMenitMasuk
	jam <- selisihMenit / 60
	menit <- selisihMenit % 60
    output(jam,menit)
endprogram