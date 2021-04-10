# Elyta Edenia
# 71180391

with open('daftarbuku.txt','r') as file:
	daftarBuku = file.read()
splitBukuWithComma = daftarBuku.split('\n')
splitTitle = []

for i in range(len(splitBukuWithComma)):
  splitTitle.append(splitBukuWithComma[i].split(','))

findTitle = input("Masukkan nama buku yang ingin dicari : ")
indexFound = -1
for i in range(len(splitTitle)):
	if findTitle in splitTitle[i][0]:
		indexFound = i

if indexFound != -1:
	ahaaa = splitTitle[indexFound]
	print(" BUKU ditemukan")
	print(" Nama: " + ahaaa[0])
	print(" Kode: " + ahaaa[1])
	print(" Tahun Rilis: " + ahaaa[2])
	print(" Deskripsi: " + ahaaa[3])
else:
	print("Barang tidak ditemukan silahkan mencari lagi")
