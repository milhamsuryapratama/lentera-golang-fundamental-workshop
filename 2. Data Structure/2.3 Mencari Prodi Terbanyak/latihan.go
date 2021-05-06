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
	if true {
		// maka tambahkan key "Informatika" pada map fakultas dengan value "Teknik Informatika", "Sistem Informatika", "RPL"
	}

	// buatlah 3 variable
	var prodi_terbanyak string
	var jumlah_prodi int
	var pertama bool

	// lakukan perulangan terhadap map fakultas menggunakan perulangan "for - range"
	for {
		// buat percabangan menggunakan "if - else", jika variable "pertama" adalah false,
		if true {
			// maka ubah value variable "pertama" menjadi true
			// ubah value jumlah_prodi menjadi panjang dari map key pada perulangan pertama
			// dan ubah value prodi_terbanyak menjadi key pertama dari map fakultas
		} else {
			// jika value dari variable "pertama" adalah true,
			// buatlah percabangan menggunakan "if - else",
			// jika value dari variable "jumlah_prodi" lebih kecil dari panjang dari key map fakultas
			if true {
				// maka ubah value variable "prodi_terbanyak" menjadi key dari map fakultas tersebut
				// dan ubah value variable "jumlah_prodi" menjadi jumlah dari value pada key map fakultas tersebut
			}
		}

		// buatlah percabangan menggunakan "if - else" dengan kondisi
		// jika panjang value dari map key fakultas == 0,
		if true {
			// maka hapus key tersebut dari map fakultas
			// untuk menghapus key yang tidak mempunyai value, kita bisa menggunakan fungsi delete
			// https://www.golangprograms.com/how-to-delete-or-remove-element-from-a-map.html
		}
	}

	// tampilkan semua daftar fakultas dan prodi terbanyak beserta jumlah prodinya menggunakan fungsi Println pada package fmt
}
