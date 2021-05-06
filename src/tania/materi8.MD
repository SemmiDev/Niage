# Ringkasan Materi 8 

## A. Operasi aritmatika:
* penjumlahan (+)
* pengurangan (-)
* perkalian   (*)
* pembagian   (/)
* modulus	  (%)
* eksponen    (a**b) bermakna = a pangkat b
* floor division / hasil pembagian dengan pembulatan kebawah


## B. Prioritas operasi, operational precedence
* ()
* exponen **
* perkalian dan teman-teman * / ** % //
* pertambahan dan pengurangan + -

## C. Contoh penggunaan: 

```python
a = 10
b = 3

a. operasi tambah +
hasil = a + b
print(a,'+',b,'=',hasil)

b. operasi pengurangan -
hasil = a - b
print(a,'-',b,'=',hasil)

c. operasi perkalian *
hasil = a * b
print(a,'*',b,'=',hasil)

d. operasi pembagian /
hasil = a / b
print(a,'/',b,'=',hasil)

e. operasi eksponen (pangkat) **
hasil = a ** b
print(a,'**',b,'=',hasil)

f. operasi modulus %
hasil = a % b
print(a,'%',b,'=',hasil)

g. operasi floor division //
hasil = a // b
print(a,'//',b,'=',hasil)

h. prioritas operasi, operational precedence
x = 3
y = 2
z = 4

hasil = x ** y * (z + x) / y - y % z // x
print(x,'**',y,'*',z,'+',x,'/',y,'-',y,'%',z,'//',x,'=',hasil)

hasil = x + y * z
print(x,'+',y,'*',z,'=',hasil)

# kurung akan mengambil langkah paling pertama
hasil = (x + y) * z 
print('(',x,'+',y,') *',z,'=',hasil)
```