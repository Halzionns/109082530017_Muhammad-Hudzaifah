package main

import "fmt"

type arrBalita [100]float64

func MinMax(arr arrBalita, n int, min, max *float64) {
	*min, *max = arr[0], arr[0]
	for i := 1; i < n; i++ {
		if arr[i] < *min { *min = arr[i] }
		if arr[i] > *max { *max = arr[i] }
	}
}

func bjir(arr arrBalita, n int) float64 {
	var total float64
	for i := 0; i < n; i++ {
		total += arr[i]
	}
	return total / float64(n)
}

func main() {
	var n int
	var data arrBalita
	var min, max float64

	fmt.Print("Masukan banyak data berat balita : ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scan(&data[i])
	}

	MinMax(data, n, &min, &max)
	
	fmt.Printf("\nBerat balita minimum: %.2f kg", min)
	fmt.Printf("\nBerat balita maksimum: %.2f kg", max)
	fmt.Printf("\nRerata berat balita: %.2f kg\n", bjir(data, n))
}