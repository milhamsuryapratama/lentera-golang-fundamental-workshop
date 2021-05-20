package main

import (
	"fmt"
)

func main() {
	bulan := map[string]string{
		"satu":     "Januari",
		"dua":      "Febaruari",
		"tiga":     "Maret",
		"empat":    "April",
		"lima":     "Mei",
		"enam":     "Juni",
		"tujuh":    "Juli",
		"delapan":  "Agustus",
		"sembilan": "September",
		"sepuluh":  "Oktober",
		"sebelas":  "November",
	}

	_, isExist := bulan["duabelas"]
	if !isExist {
		bulan["duabelas"] = "Desember"
	}

	for key, b := range bulan {
		fmt.Printf("bulan %s adalah %s \n", key, b)
	}
}
