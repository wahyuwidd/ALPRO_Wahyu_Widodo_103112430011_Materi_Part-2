soal#5 Liga bola

soal: Buatlah program yang digunakan untuk menentukan jumlah gol terbanyak dan jumlah gol tersedikit dari suatu grup liga sepak bola.

jawaban:
1. Baca input gol1, gol2, gol3, dan gol4 (jumlah gol masing-masing tim).
2. Bandingkan gol1 dan gol2:
    - Jika gol1 > gol2, maka maks1 = gol1 dan min1 = gol2.
    - Jika gol2 >= gol1, maka maks1 = gol2 dan min1 = gol1.
3. Bandingkan maks1 dan gol3.
    - Jika maks1 > gol3, maka maks2 = maks1.
    - Jika gol3 >= maks1, maka maks2 = gol3.
4. Bandingkan min1 dan gol3.
    - Jika min1 < gol3, maka min2 = min1.
    - Jika gol3 <= min1, maka min2 = gol3.
5. Bandingkan maks2 dan gol4.
    - Jika maks2 > gol4, maka maks = maks2.
    - Jika gol4 >= maks2, maka maks = gol4.
6. Bandingkan min2 dan gol4.
    - Jika min2 < gol4, maka min = min2.
    - Jika gol4 <= min2, maka min = gol4.
7. Tampilkan maks dan min.