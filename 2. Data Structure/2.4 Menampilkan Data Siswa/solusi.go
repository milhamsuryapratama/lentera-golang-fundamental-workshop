package main

import (
	"fmt"
)

type JK string

const (
	L JK = "Laki - Laki"
	P JK = "Perempuan"
	O JK = "Other"
)

type Kelas struct {
	NamaKelas string
}

type Siswa struct {
	Nama         string
	Alamat       string
	JenisKelamin JK
	Kelas
}

func main() {
	var siswa Siswa

	siswa.Nama = "Ilham"
	siswa.Alamat = "Probolinggo"
	siswa.JenisKelamin = L
	siswa.Kelas.NamaKelas = "RPL"

	fmt.Println(siswa)
}
