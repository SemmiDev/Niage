import math

a = 1
b = -5
c = 6

D = pow(b, 2) - 4 * a * c

if D < 0:
	print("Akar Imajiner")
elif D == 0:
	print("1 Akar Real")
else:
	print("2 Akar Real")


x1 = (-b + (math.sqrt(b * b - 4 * a * c))) / (2 * a)
x2 = (-b - (math.sqrt(b * b - 4 * a * c))) / (2 * a)

print(f"hasil akar (x^2 - 5x + 6) => x1 = {x1}, x2 = {x2}")