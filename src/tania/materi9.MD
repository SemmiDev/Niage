# MATERI 9

Program konversi suhu

table konversi : https://saintif.com/wp-content/uploads/2020/09/tabel-rumus-konversi-suhu-lengkap.jpg 

#Implementasi:

```python
print("\nPROGRAM KONVERSI TEMPERATUR\n")

celcius = float(input('Masukan suhu dalam celcius : '))
print("suhu adalah",celcius, "Celcius")

# reamur
reamur = (4/5) * celcius
print("suhu dalam reamur adalah ",reamur, "Reamur")

# fahrenheit
fahrenheit = ((9/5) * celcius) + 32
print("suhu dalam fahrenheit adalah ",fahrenheit, "Fahrenheit")

# Kelvin
kelvin = celcius + 273
print("suhu dalam kelvin adalah ",kelvin, "Kelvin")

```