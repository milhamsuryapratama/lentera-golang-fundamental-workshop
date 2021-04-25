package main

import "fmt"

func main() {
	var nilai = [5]int{90,80,85,95,70}

	var total_nilai int = 0

	for _, n := range nilai {
		total_nilai += n
	}

	var nilai_rata_rata = total_nilai / len(nilai)
	fmt.Println("Nilai rata rata :", nilai_rata_rata)
}
