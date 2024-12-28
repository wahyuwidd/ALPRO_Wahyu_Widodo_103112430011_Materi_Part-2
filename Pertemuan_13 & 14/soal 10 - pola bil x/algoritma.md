soal#2 Pola Bilangan X

soal: Sebuah program digunakan untuk menampilkan pola tertentu.

jawaban:
1. Baca input x (bilangan bulat positif).
2. Lakukan loop pertama dari i = 1 sampai i <= x (untuk setiap baris).
3. Lakukan loop kedua dari j = 1 sampai j <= x (untuk setiap kolom dalam baris).
4. Di dalam loop dalam, periksa dua kondisi:
    - Apakah j sama dengan i? (Ini untuk diagonal kiri ke kanan)
    - Apakah j sama dengan x - i + 1? (Ini untuk diagonal kanan ke kiri)
5. Jika salah satu dari kedua kondisi di atas benar, cetak nilai i.
6. Jika kedua kondisi salah, cetak spasi.
7. Setelah loop dalam selesai (setelah memproses semua kolom dalam satu baris), pindah ke baris baru.