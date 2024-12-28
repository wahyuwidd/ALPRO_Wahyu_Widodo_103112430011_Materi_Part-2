program Login
kamus
    jumlahGagal : integer
	username, password : char
algoritma
    jumlahGagal <- 0
    input(username, password)
    while username not "admin" OR password not "admin" do
        input(username, password)
        jumlahGagal <- jumlahGagal + 1
    endwhile
    output(jumlahGagal)
    output("Login Berhasil")
endprogram