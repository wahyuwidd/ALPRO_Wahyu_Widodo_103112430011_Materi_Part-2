soal#8 Pola Bilangan A

soal: Sebuah program digunakan untuk menampilkan pola tertentu

jawaban:
1. Baca input x
2. Lakukan loop pertama dari i = 1 sampai i <= x (untuk setiap baris)
3. Lakukan loop kedua dari j = 1 sampai j <= 2*x-1 (untuk setiap kolom dalam baris)
4. Di dalam loop kedua, periksa tiga kondisi berikut:
    - j == x-i+1: Kondisi ini membentuk garis miring kiri atas "A"
    - j == x+i-1: Kondisi ini membentuk garis miring kanan atas "A"
    - (i == x/2+1 && j > x-i+1 && j < x+i-1 && x%2 != 0): Kondisi ini membentuk garis horizontal tengah "A" Kondisi x%2 != 0 memastikan garis tengah hanya muncul jika tinggi x ganjil
5. Jika salah satu dari ketiga kondisi di atas benar, tampilkan nilai i (baris angka)
6. Jika ketiga kondisi salah, tampilkan spasi
7. Setelah loop dalam selesai (setelah memproses semua kolom dalam satu baris), pindah ke baris baru