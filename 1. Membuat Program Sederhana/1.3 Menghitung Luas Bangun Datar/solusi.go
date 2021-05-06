package main

import "fmt"

func hitungLuas(alas float32, tinggi float32) float32 {
	luas := 0.5 * alas * tinggi

	return luas
}

func main() {
	var (
		alas   float32 = 7
		tinggi float32 = 10
	)

	luas := hitungLuas(alas, tinggi)

	fmt.Println("Alas :", alas)
	fmt.Println("Tinggi :", tinggi)
	fmt.Println("Luas segitiga: ", luas)
}
