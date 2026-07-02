package main

import "fmt"

func hitungSkor(soal *int, skor *int) {
	*soal = 0
	*skor = 0
	var waktu int
	
	for i := 0; i < 8; i++ {
		fmt.Scan(&waktu)
		if waktu < 301 {
			*soal++
			*skor += waktu
		}
	}
}

func main() {
	var nama string
	var pemenangNama string
	pemenangSoal := -1
	pemenangSkor := 999999 

	for {
		fmt.Scan(&nama)
		
		if nama == "Selesai" {
			break
		}

		var soal, skor int
		hitungSkor(&soal, &skor)

		if soal > pemenangSoal {
			pemenangNama = nama
			pemenangSoal = soal
			pemenangSkor = skor
		} else if soal == pemenangSoal {
			if skor < pemenangSkor {
				pemenangNama = nama
				pemenangSoal = soal
				pemenangSkor = skor
			}
		}
	}
	fmt.Printf("%s %d %d\n", pemenangNama, pemenangSoal, pemenangSkor)
}