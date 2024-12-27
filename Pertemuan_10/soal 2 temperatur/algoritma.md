soal#2 Temperatur

soal: Sebuah sensor digunakan untuk mencatat perubahan temperatur suatu zat radioaktif X. Buatlah program yang digunakan untuk pencatatan temperatur 1  pada zat radioaktif X tersebut. 

jawaban:
1. Baca input temp1, temp2, temp3, temp4, dan temp5 (lima temperatur)
2. Inisialisasi variabel stabil dengan true
3. Inisialisasi variabel naik dan turun dengan false
4. Periksa arah perubahan antara temp1 dan temp2
    - Jika temp2 > temp1, set naik menjadi true
    - Jika temp2 < temp1, set turun menjadi true
5. Lakukan pengecekan untuk sisa pasangan temperatur (temp2 dan temp3, temp3 dan temp4, temp4 dan temp5).
    - Jika naik bernilai true dan ada pasangan temperatur yang turun (temp[i+1] < temp[i]), set stabil menjadi false.
    - Jika turun bernilai true dan ada pasangan temperatur yang naik (temp[i+1] > temp[i]), set stabil menjadi false.
6. Jika stabil bernilai true, tampilkan "Stabil naik/turun"
7. Jika stabil bernilai false, tampilkan "Tidak stabil"