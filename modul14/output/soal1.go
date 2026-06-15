package main

import (
	"fmt"
)

const NMax = 1000
type arrInt [NMax]int

func main() {
	var data arrInt
	var n, input int

	for {
		fmt.Scan(&input)
		if input < 0 {
			break
		}
		data[n] = input
		n++
	}

	insertionSort(&data, n)

	for i := 0; i < n; i++ {
		fmt.Printf("%d ", data[i])
	}
	fmt.Println()

	cekJarakData(data, n)
}

func insertionSort(T *arrInt, n int) {
	var temp, i, j int
	i = 1
	for i <= n-1 {
		j = i
		temp = T[j]
		
		for j > 0 && temp < T[j-1] {
			T[j] = T[j-1]
			j = j - 1
		}
		T[j] = temp
		i = i + 1
	}
}

func cekJarakData(T arrInt, n int) {
	if n < 2 {
		fmt.Println("Data berjarak tidak tetap")
		return
	}

	jarak := T[1] - T[0]
	isTetap := true

	for i := 1; i < n-1; i++ {
		if T[i+1]-T[i] != jarak {
			isTetap = false
			break
		}
	}

	if isTetap {
		fmt.Printf("Data berjarak %d\n", jarak)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}