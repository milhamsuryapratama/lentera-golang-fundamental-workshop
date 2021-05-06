package main

import "fmt"

func ubah_nama(nama *string) {
	*nama = "ilham"
	fmt.Println("nama dari fungsi ubah_nama :", *nama)
}

func main() {
	var nama string = "aka"

	ubah_nama(&nama)

	fmt.Println("nama dari fungsi main :", nama)
}
