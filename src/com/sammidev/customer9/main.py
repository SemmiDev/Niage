def testone():
	angka = input("masukkan angka : ")
	angka = int(angka)
	
	for i in range(angka):
		if i % 2 == 0:
			if (i % 6 == 0) or (i % 8 == 0):
				pass
			else:
				print(str(i), end = ' ')
		else:
			pass

def testtwo():
	angka = input("Mau perkalian berapa : ")
	angka = int(angka)
	
	for i in range(1, angka+1):
		print(f"{angka} x {i}  = {i * angka}")

def testthree():
	avg = 0
	temp = []
	status = True
	while(status):
		if len(temp) == 0:
			nilai = int(input("Masukkan nilai pertama: "))
			temp.append(nilai)
		nilai = int(input("Masukkan nilai selanjutnya (atau 0 untuk keluar) : "))
		if nilai == 0:
			for x in temp:
				avg = (avg + x)
			avg = avg/len(temp)
			# print(f"{angka} x {i}  = {i * angka}")
			print(f"Nilai rata-ratanya adalah {avg}")
			status = False
		else:
			temp.append(nilai)

def testfourth():
	message = """Selamat datang di McD. Berikut pilihan menunya 
  1. Ayam
  2. Burger
  3. Minuman dan Appetizer
  4. Exit
	"""
	print(message)

	choice = int(input("Masukan pilihan anda : "))
	qty = 0

	if choice == 1:
		qty = int(input("Masukkan berapa banyak potong ayam : "))
		print(f"Ayam yang anda pesan sebanyak {qty}")
	elif choice == 2:
		qty = int(input("Masukkan berapa banyak burger : "))
		print(f"burger yang anda pesan sebanyak {qty}")
	elif choice == 3:
		qty = int(input("Masukkan berapa banyak minuman dan appetizer : "))
		print(f"minuman dan appetizer yang anda pesan sebanyak {qty}")
	elif choice == 4:
		exit()
	else:
		exit()

# testone()
# testtwo()
# testthree()
testfourth()