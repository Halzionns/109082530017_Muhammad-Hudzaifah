package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Masukkan jumlah kelinci : ")
	fmt.Scan(&n)

	var beratKelinci [1000]float64

	fmt.Print("Masukkan berat beratnya : ")
	
	for i := 0; i < n; i++ {
		fmt.Scan(&beratKelinci[i])
	}

	if n > 0 {
		min := beratKelinci[0]
		max := beratKelinci[0]

		for i := 1; i < n; i++ {
			if beratKelinci[i] < min {
				min = beratKelinci[i]
			}
			if beratKelinci[i] > max {
				max = beratKelinci[i]
			}
		}
		fmt.Printf("\nTeringan : %v\n", min)
		fmt.Printf("Terberat : %v\n", max)
	}
}