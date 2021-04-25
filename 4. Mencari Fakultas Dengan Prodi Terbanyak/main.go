package main

// import package fmt
// https://golang.org/pkg/fmt

func main() {
	// buatlah sebuah map dengan nama variable fakultas
	// bertipe tipe data string dan mempunyai value slice bertipe data string
	// key 1 adalah "Pendidikan" dengan value "Pendidikan Olahraga", "Pendidikan Guru"
	// key 2 adalah "Kesehatan" dengan value "Farmasi", "Rekam Media", "Kesehatan Gigi", "Dokter Umum", "Dokter Jantung"
	// key 3 adalah "Mipa" dengan value slice kosong
	// https://blog.golang.org/maps | https://gobyexample.com/maps

	// lakukan validasi apakah terdapat key "Informatika" pada map fakultas
	// caranya adalah dengan menggunakan fungsi isExist
	// https://www.golangprograms.com/how-to-check-if-a-map-contains-a-key-in-go.html

	// lakukan percabangan, jika tidak terdapat key "Informatika" pada map fakultas
	// maka tambahkan key "Informatika" pada map fakultas dengan value "Teknik Informatika", "Sistem Informatika", "RPL"

	// buatlah 3 variable
	// variable 1 adalah "prodi_terbanyak" dengan tipe data string
	// variable 2 adalah "jumlah_prodi" dengan tipe data int
	// variable 3 adalah "pertama" dengan tipe data boolean

	// lakukan perulangan terhadap map fakultas menggunakan perulangan "for - range"

	// buat percabangan menggunakan "if - else", jika variable "pertama" adalah false,
	// maka ubah value variable "pertama" menjadi true
	// ubah value jumlah_prodi menjadi panjang dari map key pada perulangan pertama
	// dan ubah value prodi_terbanyak menjadi key pertama dari map fakultas

	// jika value dari variable "pertama" adalah true,
	// maka buatlah percabangan menggunakan "if - else" untuk memvalidasi jumlah prodi
	// buatlah percabangan menggunakan "if - else",
	// jika value dari variable "jumlah_prodi" lebih kecil dari panjang dari key map fakultas
	// maka ubah value variable "prodi_terbanyak" menjadi key dari map fakultas tersebut
	// dan ubah value variable "jumlah_prodi" menjadi jumlah dari value pada key map fakultas tersebut

	// buatlah percabangan menggunakan "if - else" dengan kondisi
	// jika panjang value dari map key fakultas == 0,
	// maka hapus key tersebut dari map fakultas
	// untuk menghapus key yang tidak mempunyai value, kita bisa menggunakan fungsi delete
	// https://www.golangprograms.com/how-to-delete-or-remove-element-from-a-map.html

	// tampilkan semua daftar fakultas dan prodi terbanyak beserta jumlah prodinya menggunakan fungsi Println pada package fmt
}
