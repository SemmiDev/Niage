# Elyta Edenia
# 71180391

def soalNo3():
	next = True
	text = """
=========TUGAS!=========
1. Tambah Data
2. Lihat Data
3. Exit
		"""
	dataNama = []
	dataPhone = [] 
	while(next):
		print(text)
		choice = int(input("masukkan nomor menu yang diinginkan : "))
		if choice == 3:
			print("berhasil logout")
			break
		elif choice == 1:
			many = int(input("masukkan jumlah alamat telepon : "))
			for x in range(many):
				nama = input("masukkan nama: ")
				telp = input("masukkan no telepon: ")
				dataNama.append(nama)
				dataPhone.append(telp)
				if x == many-1:
					print("berhasil dimasukkan")
		elif choice == 2:
			for a, b in zip(dataNama, dataPhone):
				print("Nama\t\tnomor telp")
				print(a + "\t\t" + b + "\n")

soalNo3()
