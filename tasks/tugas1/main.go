package main

import (
	"fmt"
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

var PlayList []Lagu
var durations []Waktu

func main() {
	buatPlaylist()
	cetakPlaylist()
}

// SUB PROGRAM PENGISIAN PLAYLIST
func buatPlaylist()  {
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
			saveLagu(lagu)
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

func saveLagu(l Lagu) {
	PlayList = append(PlayList, l)
}

// SUB PROGRAM UNTUK MENAMPILKAN IS PLAYLIST SESUAI KETENTUAN
func cetakPlaylist()  {
	for i , v := range PlayList {
		if i == durasiTerlama() {
			show := fmt.Sprintf("*%s %d menit %d detik",v.Judul, v.Durasi.Menit, v.Durasi.Detik )
			fmt.Println(show)
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