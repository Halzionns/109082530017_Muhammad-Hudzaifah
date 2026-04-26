package main

import (
	"fmt"
	"math"
)

func main() {
	var n, x, idxHapus, cariBilangan int
	
	fmt.Print("Masukkan jumlah elemen (N): ")
	fmt.Scan(&n)
	
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Isi elemen ke-%d: ", i)
		fmt.Scan(&arr[i])
	}

	fmt.Println("\nIsi seluruh array:", arr)

	fmt.Print("Elemen indeks ganjil: ")
	for i := 1; i < len(arr); i += 2 {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()

	fmt.Print("Elemen indeks genap: ")
	for i := 0; i < len(arr); i += 2 {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()

	fmt.Print("Masukkan nilai x untuk kelipatan indeks: ")
	fmt.Scan(&x)
	fmt.Printf("Elemen indeks kelipatan %d: ", x)
	for i := 0; i < len(arr); i++ {
		if i%x == 0 {
			fmt.Printf("%d ", arr[i])
		}
	}
	fmt.Println()

	var total float64
	for _, v := range arr {
		total += float64(v)
	}
	rataRata := total / float64(len(arr))
	
	var jumlahKuadratSelisih float64
	for _, v := range arr {
		jumlahKuadratSelisih += math.Pow(float64(v)-rataRata, 2)
	}
	stdDev := math.Sqrt(jumlahKuadratSelisih / float64(len(arr)))

	fmt.Printf("Rata-rata: %.2f\n", rataRata)
	fmt.Printf("Standar Deviasi: %.2f\n", stdDev)

	fmt.Print("Masukkan bilangan yang ingin dicari frekuensinya: ")
	fmt.Scan(&cariBilangan)
	count := 0
	for _, v := range arr {
		if v == cariBilangan {
			count++
		}
	}
	fmt.Printf("Frekuensi bilangan %d: %d kali\n", cariBilangan, count)

	fmt.Print("Masukkan indeks yang ingin dihapus: ")
	fmt.Scan(&idxHapus)
	arr = append(arr[:idxHapus], arr[idxHapus+1:]...)
	fmt.Println("   Isi array setelah dihapus:", arr)
}