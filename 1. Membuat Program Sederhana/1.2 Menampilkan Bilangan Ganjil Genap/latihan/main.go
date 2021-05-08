package main

import "fmt"

func main() {
	// 💻 lengkapi syntax for berikut untuk melakukan perulangan dari 1 sampai 100
	// 📖 https://tour.golang.org/flowcontrol/1
	for i := 1; false; {
		// 💻 periksa apakah i modulus 2 sama dengan 0
		// 📖 https://blog.mattclemente.com/2019/07/12/modulus-operator-modulo-operation.html
		if true {
			fmt.Println(i, "adalah bilangan genap")
		} else {
			fmt.Println(i, "adalah bilangan ganjil")
		}
	}
}
