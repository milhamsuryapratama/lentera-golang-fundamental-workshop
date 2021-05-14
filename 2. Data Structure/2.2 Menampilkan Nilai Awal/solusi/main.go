package main

import (
	"fmt"
)

func main() {
	nilai := [7]int{80, 90, 87, 70, 95, 84, 70}

	tiga_nilai_awal := nilai[:3]

	tiga_nilai_awal = append(tiga_nilai_awal, 70)

	fmt.Println(tiga_nilai_awal)

	fmt.Println("kapasitas slice tiga_nilai_awal", cap(tiga_nilai_awal))
}
