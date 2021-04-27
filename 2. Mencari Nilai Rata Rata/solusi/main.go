package main

import "fmt"

func main() {
	var nilai = [5]int{90, 80, 85, 95, 70}

	var total_nilai int = 0

	for _, n := range nilai {
		var nilaiElement = n
		total_nilai += nilaiElement
	}

	var jumlah_nilai int = len(nilai)

	var nilai_rata_rata int = total_nilai / jumlah_nilai
	fmt.Println("Nilai rata rata :", nilai_rata_rata)
}
