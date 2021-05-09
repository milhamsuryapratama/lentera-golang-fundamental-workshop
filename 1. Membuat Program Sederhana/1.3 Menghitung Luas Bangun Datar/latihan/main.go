package main

import "fmt"

// 💻 buatlah sebuah fungsi dengan nama hitungLuas
// 💻 berikan parameter alas dan tinggi dengan tipe data float32
// 💻 pastikan fungsi hitungLuas mereturn tipe data float32
// 📖 https://tour.golang.org/basics/4
func hitungLuas(alas, tinggi float32) (float32, string) {
	// catatan : kode didalam blok function
	// 💻 buatlah varibale luas dan hitung luas segitiga dengan rumus 0.5 * alas * tinggi
	luas := 0.5 * alas * tinggi

	// 💻 return nilai dari variable luas
	return luas, "Luas Segitiga"
}

func main() {
	// 💻 buatlah variable alas dan tinggi dengan tipe data float32 dan isi value dari kedua variable dengan angka bebas
	var alas float32 = 7
	var tinggi float32 = 10

	// 💻 panggil fungsi hitungLuas dan isi parameternya dengan variable alas dan tinggi
	// 💻 simpan hasil kembalian (return) dari fungsi hitungLuas pada variable hasilLuas
	hasilLuas, hasil := hitungLuas(alas, tinggi)

	// tampilkan variable hasilLuas dengan fungsi println
	fmt.Println(hasil, hasilLuas)
}
