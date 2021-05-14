package main

import (
	"fmt"
)

type Siswa struct {
	Nama   string
	Alamat string
}

func (s *Siswa) SetNama() {
	s.Nama = "ilham"
}

func (s *Siswa) SetAlamat() {
	s.Alamat = "Probolinggo"
}

func (s *Siswa) GetNama() string {
	return s.Nama
}

func (s *Siswa) GetAlamat() string {
	return s.Alamat
}

func main() {
	var s Siswa

	s.SetNama()
	s.SetAlamat()

	nama := s.GetNama()
	alamat := s.GetAlamat()

	fmt.Println(nama)
	fmt.Println(alamat)
}
