soal#1 Prima

soal: Sebuah bilangan prima adalah bilangan yang hanya memiliki dua faktor, yaitu 1 dan bilangan itu sendiri

jawaban:
1. Baca input x (bilangan bulat positif).
2. Jika x kurang dari atau sama dengan 1, kembalikan false (karena 1 dan bilangan negatif bukan prima).
3. Lakukan loop dari i = 2 sampai i*i <= x. Kita hanya perlu memeriksa pembagi hingga akar kuadrat dari x untuk menentukan apakah bilangan tersebut prima.
4. Di dalam loop, periksa apakah x habis dibagi i (x % i == 0). Jika ya, kembalikan false (karena x memiliki faktor selain 1 dan dirinya sendiri).
5. Jika loop selesai tanpa menemukan faktor selain 1 dan x, kembalikan true (karena x adalah bilangan prima).