package main

import (
	"fmt"
	"os"
)

func main(){
	testfourth()
}

func testone()  {
	var a int

	fmt.Print("Masukkan angka : ")
	fmt.Scanln(&a)

	for i := 2; i < a ; i++ {
		if i % 2 == 0 {
			if i % 6 == 0 || i % 8 == 0{
			}else {
				fmt.Print(i, " ")
			}
		}
	}
}
func testtwo()  {
	var a int

	fmt.Print("Mau perkalian berapa : ")
	fmt.Scanln(&a)

	for i := 1; i <= 10 ; i++ {
		fmt.Printf("%d x %d = %d \n", a,i, a*i)
	}
}
func tessthree() {
	var nilai []float64
	for {
		var temp float64
		if len(nilai)  == 0{
			fmt.Print("Masukkan nilai pertama: ")
			fmt.Scan(&temp)
			nilai = append(nilai, temp)
		}
		fmt.Print("Masukkan nilai selanjutnya (atau 0 untuk keluar): ")
		fmt.Scan(&temp)
		if temp == 0 {
			var result float64
			for i := 0; i < len(nilai); i++ {
				result += nilai[i]
			}
			fmt.Printf("Nilai rata-ratanya adalah: %f", result / (float64(len(nilai))))
			os.Exit(0)
		}
		nilai = append(nilai, temp)
	}
}

func testfourth() {
	message := "Selamat datag di McD. Berikut pilihan menunya\n" +
		"1. Ayam\n" +
		"2. Burger\n" +
		"3. Minuman dan Appetizer\n" +
		"4. Exit"
	for {
		var pilihan, qty int
		fmt.Println(message)
		fmt.Print("Masukkan pilihan anda : ")
		fmt.Scan(&pilihan)
		if pilihan == 1 {
			fmt.Println("Masukkan berapa banyak potong ayam : ")
			fmt.Scan(&qty)
			fmt.Println("Ayam yang anda pesan sebanyak ", qty)
		}else if pilihan == 2 {
			fmt.Println("Masukkan berapa banyak burger : ")
			fmt.Scan(&qty)
			fmt.Println("burger yang anda pesan sebanyak ", qty)
		}else if pilihan == 3 {
			fmt.Println("Masukkan berapa banyak minuman dan appetizer : ")
			fmt.Scan(&qty)
			fmt.Println("minuman dan appetizer yang anda pesan sebanyak ", qty)
		}else if pilihan  == 4{
			os.Exit(0)
		}else {
			fmt.Println("Inputa salah")
		}
	}

}