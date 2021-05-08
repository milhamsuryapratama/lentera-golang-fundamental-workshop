package main

// 💻 import package `fmt` untuk dapat menggunakan fungsi menampilkan data
// 📖 https://golang.org/pkg/fmt
import "fmt"

func main() {
	// 💻 buatlah sebuah variable dengan nama nilai, bertipe data int dengan value 95
	// 📖 https://gobyexample.com/variables
	//var nilai int = 95
	nilai := 95

	// ini referensi tentang if - else, mungkin bisa membantu untuk memahami kode dibawah
	// 📖 https://gobyexample.com/if-else

	// ubah kondisi dibawah ini, jika nilai lebih dari sama dengan 90
	// 📖 https://golang.org/ref/spec#Operators
	if nilai >= 90 {
		// maka tampilkan `nilai anda sempurna, terus belajar` pada layar
		// 📖 https://golang.org/pkg/fmt/#Println
		fmt.Println("nilai anda sempurna, terus belajar")
		//jika nilai lebih besar sama dengan 75
	} else if nilai >= 75 {
		// maka tampilkan `nilai anda bagus, terus tingkatkan` pada layar
		// 📖 https://golang.org/pkg/fmt/#Println
		fmt.Println("nilai anda bagus, terus tingkatkan")
	} else {
		// maka tampilkan `nilai anda kurang, terus semangat` pada layar
		// 📖 https://golang.org/pkg/fmt/#Println
		fmt.Println("nilai anda kurang, terus semangat")
	}
}
