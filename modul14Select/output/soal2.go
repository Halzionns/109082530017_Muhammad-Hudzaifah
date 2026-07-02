package main

import "fmt"

func sortAsc(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

func sortDesc(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] > arr[maxIdx] {
				maxIdx = j
			}
		}
		arr[i], arr[maxIdx] = arr[maxIdx], arr[i]
	}
}

func main() {
	var n int
	
	fmt.Scan(&n)
	
	for i := 0; i < n; i++ {
		var m int
		fmt.Scan(&m)
		
		var ganjil []int
		var genap []int
		
		for j := 0; j < m; j++ {
			var num int
			fmt.Scan(&num)
			
			if num%2 != 0 {
				ganjil = append(ganjil, num)
			} else { 
				genap = append(genap, num)
			}
		}
		
		sortAsc(ganjil)
		sortDesc(genap)
		
		for j := 0; j < len(ganjil); j++ {
			fmt.Printf("%d ", ganjil[j])
		}
		
		for j := 0; j < len(genap); j++ {
			fmt.Printf("%d ", genap[j])
		}
		fmt.Println()
	}
}