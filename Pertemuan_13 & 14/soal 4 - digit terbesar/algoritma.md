soal#4 Digit Terbesar

soal: Sebuah program digunakan untuk mencari digit terbesar dari suatu bilangan bulat positif

jawaban:
1. Baca input bilangan (bilangan bulat positif)
2. Inisialisasi variabel digitTerbesar dengan 0 (karena digit terkecil adalah 0)
3. Lakukan loop selama bilangan lebih besar dari 0:
    - Ambil digit terakhir dari bilangan menggunakan operator modulo 10: digit = bilangan % 10
    - Jika digit lebih besar dari digitTerbesar, maka update digitTerbesar dengan nilai digit
    - Hilangkan digit terakhir dari bilangan dengan pembagian integer 10: bilangan = bilangan / 10
4. Setelah loop selesai, tampilkan nilai digitTerbesar