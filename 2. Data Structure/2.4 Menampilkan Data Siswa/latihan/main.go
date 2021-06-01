package main

import "fmt"

// 💻 buatlah sebuah type dengan nama JK dan bertipe data string

// 💻 buatlah sebuah konstanta dan isi dengan
// 1. L bertipe data JK dengan value Laki - Laki
// 2. P bertipe data JK dengan value Perempuan
// 3. O bertipe data JK dengan value Other
// 📖 https://golang.org/doc/effective_go#constants

// 💻 buatlah sebuah struct dengan nama Kelas, isi dengan properti
// NamaKelas bertipe data string
// lihat cara membuat struct pada kode struct `Vertex` setelah `import "fmt"` (referensi ke-1)
// 📖 https://tour.golang.org/moretypes/2 | https://gobyexample.com/structs

// 💻 buatlah sebuah struct dengan nama Siswa, isi dengan properti
// Nama bertipe data string
// Alamat bertipe data string
// JenisKelamin bertipe data JK
// 💻 dan buatlah embedded struct Kelas pada struct pada struct Siswa
// kode untuk membuat embedded struct bisa dilihat pada poin A.24.5
// 📖 https://dasarpemrogramangolang.novalagung.com/A-struct.html

func main() {
	// 💻 inisialiasi sebuah variable dengan nama siswa bertipe data struct Siswa
	// inisialisasi variable dengan tipe data struct bisa dilihat pada point A.24.2
	// 📖 https://dasarpemrogramangolang.novalagung.com/A-struct.html

	// 💻 isi properti Nama, Alamat, JenisKelamin dan NamaKelas
	// untuk mengiri value dari properti struct bisa dilihat pada point A.24.3
	// 📖 https://dasarpemrogramangolang.novalagung.com/A-struct.html

	// tampilkan variable siswa menggunakan fungsi Println pada package fmt
	fmt.Println()
}
