soal#4 Parkir

soal: Sebuah tempat wisata buka jam 7:00 pagi hingga 5:00 sore (17:00 dalam format 24 jam). Buatlah program yang digunakan untuk menghitung durasi parkir suatu kendaraan pada tempat wisata tersebut.

jawaban:
1. Baca input h1, m1, h2, dan m2
2. Konversi waktu ke menit:
    - total_menit_masuk = h1 * 60 + m1
    - total_menit_keluar = h2 * 60 + m2
3. Jika h2 < h1, tambahkan 12 jam (720 menit) ke total_menit_keluar untuk menangani kasus parkir melewati tengah hari
4. Hitung selisih menit: selisih_menit = total_menit_keluar - total_menit_masuk
5. Hitung jam dan menit dari selisih menit:
    - jam = selisih_menit / 60 (pembagian integer)
    - menit = selisih_menit % 60 (modulo)
6. Tampilkan jam dan menit