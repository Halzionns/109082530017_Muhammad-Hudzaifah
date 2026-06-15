package main

import (
	"fmt"
)

const nMax = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

type DaftarBuku [nMax]Buku

func main() {
	var pustaka DaftarBuku
	var n, r int

	DaftarkanBuku(&pustaka, &n)

	fmt.Println("=== Buku Terfavorit ===")
	CetakTerfavorit(pustaka, n)

	UrutBuku(&pustaka, n)

	fmt.Println("=== 5 Buku Rating Tertinggi ===")
	Cetak5Terbaru(pustaka, n)

	fmt.Scan(&r) 
	fmt.Printf("=== Hasil Pencarian Rating %d ===\n", r)
	CariBuku(pustaka, n, r)
}

func DaftarkanBuku(pustaka *DaftarBuku, n *int) {
	fmt.Scan(n)
	for i := 0; i < *n; i++ {
		fmt.Scan(&pustaka[i].id, &pustaka[i].judul, &pustaka[i].penulis, &pustaka[i].penerbit, 
			&pustaka[i].eksemplar, &pustaka[i].tahun, &pustaka[i].rating)
	}
}

func CetakTerfavorit(pustaka DaftarBuku, n int) {
	if n == 0 {
		return
	}
	
	maxIdx := 0
	for i := 1; i < n; i++ {
		if pustaka[i].rating > pustaka[maxIdx].rating {
			maxIdx = i
		}
	}
	
	b := pustaka[maxIdx]
	fmt.Printf("%s, %s, %s, %d\n", b.judul, b.penulis, b.penerbit, b.tahun)
}

func UrutBuku(pustaka *DaftarBuku, n int) {
	var temp Buku
	var i, j int

	for i = 1; i < n; i++ {
		temp = pustaka[i]
		j = i
		for j > 0 && temp.rating > pustaka[j-1].rating {
			pustaka[j] = pustaka[j-1]
			j = j - 1
		}
		pustaka[j] = temp
	}
}

func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	limit := 5
	if n < 5 {
		limit = n
	}
	
	for i := 0; i < limit; i++ {
		fmt.Printf("%d. %s\n", i+1, pustaka[i].judul)
	}
}

func CariBuku(pustaka DaftarBuku, n int, r int) {
	left := 0
	right := n - 1
	foundIdx := -1

	for left <= right && foundIdx == -1 {
		mid := (left + right) / 2
		
		if pustaka[mid].rating == r {
			foundIdx = mid
		} else if pustaka[mid].rating < r {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	if foundIdx != -1 {
		b := pustaka[foundIdx]
		fmt.Printf("%s, %s, %s, %d, %d, %d\n", 
			b.judul, b.penulis, b.penerbit, b.tahun, b.eksemplar, b.rating)
	} else {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	}
}