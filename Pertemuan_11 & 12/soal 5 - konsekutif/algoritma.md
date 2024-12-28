soal#5 Konsekutif

soal: Apabila didefinisikan sebuah bilangan konsekutif adalah bilangan yang selisih setiap digit yang bersebelahannya adalah satu. Maka buatlah algoritma untuk menentukan suatu bilangan konsekutif atau tidak

jawaban:
1. Baca input n (bilangan bulat positif)
2. Ambil digit terakhir dari n dan simpan di digitTerakhir
3. Hapus digit terakhir dari n
4. Inisialisasi konsekutif dengan true
5. Lakukan loop selama n masih lebih besar dari 0 (masih ada digit yang perlu dicek)
6. Ambil digit terakhir dari n dan simpan di digitSekarang
7. Periksa apakah selisih antara digitSekarang dan digitTerakhir bukan 1 atau -1 Jika bukan, set konsekutif menjadi false
8. Update digitTerakhir dengan digitSekarang
9. Hapus digit terakhir dari n
10. Setelah loop selesai, tampilkan nilai konsekutif