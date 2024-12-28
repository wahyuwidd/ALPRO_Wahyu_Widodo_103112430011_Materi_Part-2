soal#5 Cari Digit

soal: Sebuah program digunakan untuk mencari apakah suatu digit tertentu (x) terdapat di dalam suatu bilangan bulat positif (n), di mana nilai x ≤ 9 (karena x adalah digit)

jawaban:
1. Baca input x (digit yang dicari) dan n (bilangan yang akan dicari)
2. Lakukan loop selama n lebih besar dari 0:
    - Ambil digit terakhir dari n menggunakan operator modulo 10: digit = n % 10
    - Jika digit sama dengan x, maka kembalikan true (digit ditemukan)
    - Hilangkan digit terakhir dari n dengan pembagian integer 10: n = n / 10
3. Jika loop selesai tanpa menemukan x, maka kembalikan false (digit tidak ditemukan)