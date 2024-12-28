soal#7 Temperatur

soal: Sebuah sensor temperatur perlu dipasang di suatu lab rahasia. Sensor ini merekam perubahan suhu dan data statistik lainnya

jawaban:
1. Baca input suhu
2. Lakukan loop selama nolBerturutTurut tidak sama dengan 2:
    - Jika suhu sama dengan 0, increment nolBerturutTurut
    - Jika suhu lebih kecil dari min, update min dengan suhu
    - Jika suhu lebih besar dari max, update max dengan suhu
    - Tambahkan suhu ke total
    - Increment jumlahInput
3. Baca input suhu berikutnya
4. Kurangi 1 dari jumlahInput dan total untuk mengabaikan nol terakhir
5. Tampilkan max, min, dan total / jumlahInput (sebagai float) untuk menghitung rata rata