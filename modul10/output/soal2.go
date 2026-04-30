package main

import "fmt"

func main() {
	var x, y int

	fmt.Print("Masukkan jumlah ikan dan kapasitas tampung ikan pada setiap wadah : ")
	fmt.Scan(&x, &y)

	var beratIkan [1000]float64

	fmt.Printf("Masukkan %d data berat ikan: \n", x)
	for i := 0; i < x; i++ {
		fmt.Printf("Berat ikan ke-%d: ", i+1)
		fmt.Scan(&beratIkan[i])
	}

	var totalBeratKeseluruhan float64
	var jumlahWadah int
	SementaraWadahBerat := 0.0

	fmt.Print("Total berat per wadah: ")
	for i := 0; i < x; i++ {
		SementaraWadahBerat += beratIkan[i]
		totalBeratKeseluruhan += beratIkan[i]

		if (i+1)%y == 0 || i == x-1 {
			fmt.Printf("%.2f ", SementaraWadahBerat)
			SementaraWadahBerat = 0
			jumlahWadah++
		}
	}
	fmt.Println()

	if jumlahWadah > 0 {
		rataRata := totalBeratKeseluruhan / float64(jumlahWadah)
		fmt.Printf("Rata-rata berat per wadah: %.2f\n", rataRata)
	}
}