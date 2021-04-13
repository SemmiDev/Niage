from math import sqrt, cos

# b = 3, c = 4, alpha = 1.2

b = int(input("input b : "))
c = int(input("input c : "))
alpha = float(input("input alpha : "))

# formula -> a^2 = b^2 + c^2 - 2bc cos alpha

panjangA = (b * b) + (c * c) -  (2 * b * c) * cos(alpha)
print("hasil = ", panjangA)
