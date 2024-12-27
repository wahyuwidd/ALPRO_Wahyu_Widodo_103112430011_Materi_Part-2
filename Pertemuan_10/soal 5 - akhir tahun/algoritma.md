soal#5 Akhir tahun

soal: Sebuah minimarket sedang merayakan acara akhir tahun, sehingga diberikanlah promo pada hari tersebut berupa cashback, diskon, dan juga kartu membership. Pembeli memperoleh kartu apabila bersedia dibuatkan dan memperoleh diskon. Diskon 10% diberikan apabila belanja minimal Rp. 100.000, dan Cashback Rp. 75.000 diberikan apabila belanja minimal Rp. 200.

jawaban:
1. Baca input total_belanja (integer) dan bersedia_kartu (boolean)
2. Inisialisasi variabel boolean dapat_kartu, dapat_diskon, dan dapat_cashback
3. Tentukan apakah pembeli mendapat kartu: dapat_kartu = bersedia_kartu
4. Tentukan apakah pembeli mendapat diskon: dapat_diskon = total_belanja >= 100000
5. Tentukan apakah pembeli mendapat cashback: dapat_cashback = total_belanja >= 200000 AND dapat_kartu
6. Hitung harga yang harus dibayar:
    - Jika dapat_diskon, hitung diskon: diskon = total_belanja * 0.10
    - Kurangi diskon dari total belanja: harga_setelah_diskon = total_belanja - diskon
    - Jika dapat_cashback, kurangi cashback: harga_bayar = harga_setelah_diskon - 75000
    - Jika tidak ada diskon dan cashback, harga_bayar = total_belanja
7. Tampilkan dapat_kartu, dapat_diskon, dapat_cashback, dan harga_bayar