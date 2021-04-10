# Elyta Edenia
# 71180391

def soalNo1():
	# input = 6e616d615361796154696e61
	# hasil = 34161133418139601367651151457

	password = input("Masukkan password yang akan di generate : ")
	decimal = int(password, 16)
	r = open("passTina", "x")
	f = open("passTina", "w")
	f.write(str(decimal))
	f.close()

soalNo1()
