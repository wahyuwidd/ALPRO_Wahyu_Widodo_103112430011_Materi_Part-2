program 
kamus
    jumlahGagal : integer
	username, password : char
algoritma
    jumlahGagal = 0
    input(username, password)
    while username != "admin" OR password != "admin" do
        input(username, password)
        jumlahGagal+1
    endwhile
    output(jumlahGagal)
    output("Login Berhasil")
endprogram