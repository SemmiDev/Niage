package main

import (
	"fmt"
	"strconv"
)

const NMAX int = 1000

type (
	Waktu struct {
		Menit 	int
		Detik 	int
	}

	Lagu struct {
		Judul    string
		Penyanyi string
		Durasi   Waktu
	}
)

var durations []Waktu
var PlayList []Lagu

// SUB PROGRAM PENGISIAN PLAYLIST
func inputLagu()  {
	var judul, penyanyi string
	var menit, detik int

	for i := 1; i <= NMAX; i++ {

		fmt.Scan(&judul)
		fmt.Scan(&penyanyi)

		// cek input judul dan penyanyi
		if judul == "#" && penyanyi == "#"{
			return
		}

		fmt.Scan(&menit)
		fmt.Scan(&detik)

		waktu := Waktu{
			Menit: menit,
			Detik: detik,
		}

		lagu := Lagu{
			Judul:    judul,
			Penyanyi: penyanyi,
			Durasi:   waktu,
		}

		isExists := findDuplikasi(lagu)

		if isExists {} else {
			durations = append(durations, waktu)

			// BUAT PLAYLIST
			buatPlaylist(lagu)
		}
	}
}

func findDuplikasi(lagu Lagu) bool {
	for _, v := range PlayList {
		if (v.Penyanyi == lagu.Penyanyi) && (v.Judul == lagu.Judul) {
			return true
		}
	}
	return false
}

func buatPlaylist(l Lagu) {
	PlayList = append(PlayList, l)
}

// SUB PROGRAM UNTUK MENAMPILKAN IS PLAYLIST SESUAI KETENTUAN
func cetakPlaylist()  {
	for i , v := range PlayList {
		if i == durasiTerlama() {
			fmt.Println("*" + v.Judul + " " +  strconv.Itoa(v.Durasi.Menit) +  " menit " + strconv.Itoa(v.Durasi.Detik) +  " detik")
			continue
		}
		fmt.Println(v.Judul)
	}
}

// SUB PROGRAM UNTUK MENCARI DURASI TERLAMA
func durasiTerlama() int {

	maxDuration := convertToSeconds(durations[0])
	index := 0

	for i, v := range durations {
		if convertToSeconds(v) > maxDuration {
			maxDuration = convertToSeconds(v)
			index = i
		}
	}
	return index
}

// UTILITAS MENIT TO DETIK
func convertToSeconds(waktu Waktu) int {
	return (waktu.Menit * 60) + waktu.Detik
}

func main() {
	inputLagu()
	cetakPlaylist()
}