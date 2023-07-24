package main

import (
	"fmt"
	"os"
	"strconv"
)

type Person struct {
	Absen     int
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

type Persona struct {
	Absen     int
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func main() {

	golangForWoman := map[string]Person{
		"anita":  {Absen: 1, Nama: "Anita", Alamat: "Jakarta", Pekerjaan: "BE", Alasan: "Keperluan di pekerjaan"},
		"boonie": {Absen: 2, Nama: "Boonie", Alamat: "Bandung", Pekerjaan: "DS", Alasan: "Switch karir"},
		"clara":  {Absen: 3, Nama: "Clara", Alamat: "Semarang", Pekerjaan: "Kuliah", Alasan: "Ingin mendalami Golang"},
		"desy":   {Absen: 4, Nama: "Desy", Alamat: "Yogyakarta", Pekerjaan: "IRT", Alasan: "Mau berkarier di IT"},
	}

	golangForWomanA := map[int]Persona{
		1: {Absen: 1, Nama: "Anita", Alamat: "Jakarta", Pekerjaan: "BE", Alasan: "Keperluan di pekerjaan"},
		2: {Absen: 2, Nama: "Boonie", Alamat: "Bandung", Pekerjaan: "DS", Alasan: "Switch karir"},
		3: {Absen: 3, Nama: "Clara", Alamat: "Semarang", Pekerjaan: "Kuliah", Alasan: "Ingin mendalami Golang"},
		4: {Absen: 4, Nama: "Desy", Alamat: "Yogyakarta", Pekerjaan: "IRT", Alasan: "Mau berkarier di IT"},
	}

	if len(os.Args) < 2 {
		fmt.Println("Masukkan nama atau absen, contoh: go run main.go anita atau go run main.go 1")
		os.Exit(1)
	}
	gfwNama := os.Args[1]

	gfwAbsen := os.Args[1]
	absen, err := strconv.Atoi(gfwAbsen)
	if err != nil {
	}

	if persona, found := golangForWomanA[absen]; found {
		fmt.Printf("Absen: %d\nNama: %s\nAlamat: %s\nPekerjaan: %s\nAlasan: %s\n", persona.Absen, persona.Nama, persona.Alamat, persona.Pekerjaan, persona.Alasan)
	} else if person, found := golangForWoman[gfwNama]; found {
		fmt.Printf("Absen: %d\nNama: %s\nAlamat: %s\nPekerjaan: %s\nAlasan: %s\n", person.Absen, person.Nama, person.Alamat, person.Pekerjaan, person.Alasan)
	} else {
		fmt.Printf("Nama atau absen tidak ditemukan")
	}
}
