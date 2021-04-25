package main

import "fmt"

const phi = 3.14

type Lingkaran struct {
	JariJari float32
}

func hitungLuasLingkaran(lingkaran1 Lingkaran, lingkaran2 Lingkaran) (float32, float32) {
	luasLingkaran1 := phi * lingkaran1.JariJari * lingkaran1.JariJari

	luasLingkaran2 := phi * lingkaran2.JariJari * lingkaran2.JariJari

	return luasLingkaran1, luasLingkaran2
}

func main() {
	var lingkaran1 Lingkaran
	lingkaran1.JariJari = 14

	var lingkaran2 Lingkaran
	lingkaran2.JariJari = 24

	luasLingkaran1, luasLingkaran2 := hitungLuasLingkaran(lingkaran1, lingkaran2)

	fmt.Println("Luas Lingkaran 1", luasLingkaran1)
	fmt.Println("Luas Lingkaran 2", luasLingkaran2)

	if luasLingkaran1 > luasLingkaran2 {
		fmt.Println("Luas lingkaran1 lebih luas daripada lingkaran2")
	} else {
		fmt.Println("Luas lingkaran2 lebih luas daripada lingkaran1")

	}
}
