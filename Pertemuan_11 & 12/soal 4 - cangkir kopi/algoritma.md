soal#4 Cangkir Kopi

soal: Buatlah sebuah program untuk menghitung banyak cangkir kopi yang bisa dibuat apabila tersedia n gula dan m kopi. Satu cangkir kopi memerlukan x gula dan y kopi

jawaban:
1. Baca input n (jumlah gula), m (jumlah kopi), x (gula per cangkir), dan y (kopi per cangkir)
2. Inisialisasi jumlah_cangkir dengan 0
3. Lakukan loop selama masih ada cukup gula dan kopi untuk membuat setidaknya satu cangkir lagi:
    - Periksa apakah n >= x dan m >= y
        - Jika ya:
            - Kurangi n dengan x dan m dengan y
            - Inkrement jumlah_cangkir
        - Jika tidak:
            - loop berhenti
4. Tampilkan jumlah_cangkir