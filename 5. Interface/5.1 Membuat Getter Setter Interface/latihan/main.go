package main

import "fmt"

// 💻 buatlah sebuah interface dengan nama SiswaInterface
// 💻 Tuliskan method apa saja yang terdapat pada struct Siswa
// 📖 https://tour.golang.org/methods/9

type Siswa struct {
	Nama   string
	Alamat string
}

func (s *Siswa) SetNama() {
	s.Nama = "Ilham"
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
	// 💻 buatlah variable dengan nama "s" dan bertipe data SiswaInterface dan value pointer (ampersand) struct Siswa

	s.SetNama()
	s.SetAlamat()

	nama := s.GetNama()
	alamat := s.GetAlamat()

	fmt.Println(nama)
	fmt.Println(alamat)
}
