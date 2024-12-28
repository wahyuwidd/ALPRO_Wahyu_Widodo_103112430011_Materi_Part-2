soal#2 Bilangan Biner

soal: Buatlah program yang digunakan untuk konversi bilangan desimal (basis 10) ke bilangan biner (basis 2)

jawaban:
1. Baca input desimal (bilangan bulat desimal yang akan dikonversi)
2. Jika desimal adalah 0, maka biner adalah "0". Selesai.
3. Inisialisasi string kosong biner untuk menyimpan hasil konversi.
4. Lakukan loop selama desimal lebih besar dari 0:
    - Hitung sisa bagi sisa = desimal % 2.
    - Tambahkan sisa (dikonversi ke string) ke awal string biner.
    - Lakukan pembagian integer desimal = desimal / 2.
5. Setelah loop selesai, kembalikan string biner.