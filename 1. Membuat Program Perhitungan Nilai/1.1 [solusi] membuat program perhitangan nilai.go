package main

import "fmt"

func main() {
	var nilai int = 95

	if nilai >= 90 {
		fmt.Println("nilai anda sempurna, terus belajar")
	} else if nilai >= 75 && nilai < 90 {
		fmt.Println("nilai anda bagus, terus tingkatkan")
	} else {
		fmt.Println("nilai anda kurang, terus semangat")
	}
}
