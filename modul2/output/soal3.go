package main

import "fmt"

func main() {
	var beratTotal, kg, sisa, biayaKg, biayaSisa, totalBiaya int

	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&beratTotal)

	kg = beratTotal / 1000
	sisa = beratTotal % 1000

	biayaKg = kg * 10000

	if kg >= 10 {
		biayaSisa = 0
	} else {
		if sisa >= 500 {
			biayaSisa = sisa * 5
		} else {
			biayaSisa = sisa * 15
		}
	}

	totalBiaya = biayaKg + biayaSisa

	fmt.Printf("Detail berat: %d kg + %d gr\n", kg, sisa)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaKg, biayaSisa)
	fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)
}