package main

import "fmt"

func main() {
	var nilai int = 10
	var pNilai *int = &nilai

	fmt.Println("Nilai", nilai)   // 10
	fmt.Println("pNilai", pNilai) // 0xc0000140e0

	nilai = 5

	fmt.Println("Nilai", nilai)    // 5
	fmt.Println("pNilai", *pNilai) // 0xc0000140e0
}
