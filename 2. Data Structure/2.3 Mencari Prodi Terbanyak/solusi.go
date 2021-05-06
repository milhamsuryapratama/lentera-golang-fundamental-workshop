package main

import (
	"fmt"
)

func main() {
	fakultas := map[string][]string{
		"Pendidikan": {"Pendidikan Olahraga", "Pendidikan Guru"},
		"Kesehatan":  {"Farmasi", "Rekam Media", "Kesehatan Gigi", "Dokter Umum", "Dokter Jantung"},
		"Mipa":       {},
	}

	var _, isExist = fakultas["Informatika"]

	if isExist == false {
		fakultas["Informatika"] = []string{"Teknik Informatika", "Sistem Informatika", "RPL"}
	}

	prodi_terbanyak := ""
	jumlah_prodi := 0
	pertama := false

	for i, f := range fakultas {
		if pertama == false {
			pertama = true
			jumlah_prodi = len(f)
			prodi_terbanyak = i
		} else {
			if jumlah_prodi < len(f) {
				prodi_terbanyak = i
				jumlah_prodi = len(f)
			}
		}

		if len(f) == 0 {
			delete(fakultas, i)
		}
	}

	fmt.Println(fakultas)
	fmt.Println("prodi terbanyak adalah", prodi_terbanyak, "dengan jumlah prodi", jumlah_prodi)
}
