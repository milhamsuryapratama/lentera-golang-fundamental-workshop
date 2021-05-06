package main

import (
	"fmt"
	"sort"
)

func main() {
	nilai := [7]int{80, 90, 87, 70, 95, 84, 70}

	tiga_nilai_awal := nilai[:3]
	tiga_nilai_akhir := nilai[4:]

	sort.Ints(tiga_nilai_awal)
	sort.Ints(tiga_nilai_akhir)

	tiga_nilai_awal = append(tiga_nilai_awal, 70)
	tiga_nilai_akhir = append(tiga_nilai_akhir, 80)

	fmt.Println(tiga_nilai_awal)
	fmt.Println(tiga_nilai_akhir)

	fmt.Println("kapasitas slice tiga_nilai_awal", cap(tiga_nilai_awal))
	fmt.Println("kapasitas slice tiga_nilai_akhir", cap(tiga_nilai_akhir))
}
