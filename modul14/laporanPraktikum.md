# <h1 align="center">Laporan Praktikum Modul 14 - ... </h1>
<p align="center">[Muhammad Hudzaifah] - [109082530017]</p>

## SOAL

### 1. [Soal 1]
#### soal1.go

```go
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
```
### Output :

##### Output 
![Screenshot Output soal 1](https://github.com/Halzionns/109082530017_Muhammad-Hudzaifah/blob/main/modul14/output/output1.png)
[penjelasan] Pertama-tama, kita bikin array dengan kapasitas seribu, terus pakai looping buat ngescan masukan. Kita input terus angkanya, dan pas program ngebaca angka negatif, dia langsung ngebreak biar yang kesimpan cuma angka positif aja. Selanjutnya array dilempar ke prosedur insertionSort, Kita ambil satu angka, lalu bandingin sama angka angka di sebelah kirinya. Kalau angka di kiri lebih gede, geser ke kanan, lalu selipin di posisi yang pas. Setelah datanya rapi berurutan, angkanya tinggal di print pakai perulangan biasa. Terakhir, masuk ke cekJarakData. Kita ambil selisih dua angka pertama sebagai patokan awal, lalu program bakal looping ngecheck sisa angkanya kalau ketemu satu aja jarak antar angka yang beda, statusnya langsung berubah jadi false dan dibreak. Output akhirnya tinggal diPrintf pakai %v kalau jaraknya tetap, atau cetak "Data berjarak tidak tetap" kalau ada yang melenceng


## SOAL

### 2. [Soal 2]
#### soal2.go

```go
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
```
### Output :

##### Output 
![Screenshot Output Soal 2](https://github.com/Halzionns/109082530017_Muhammad-Hudzaifah/blob/main/modul14/output/output2.png)
[penjelasan] Pertama, kita input jumlah buku perpustakaan di Scan(&n), terus ngebuat Array pustaka pakai struct buat nampung detail data data bukunya, kapasitasnya nMax 7919. Terus di dalam looping nya, kita masukin atribut bukunya satu satu mulai dari id, judul, sampai ratingnya , dan abis itu program bakal nyari buku terfavorit pakai pencarian beruntun buat ngecek mana rating yang paling gede buat diPrint duluan. Selanjutnya, semua data buku tadi bakal diurutin dari rating tertinggi ke terendah pakai Insertion Sort, pokoknya buku yang ratingnya lebih gede bakal terus digeser ke posisi depan. Kalo udah terurut, 5 judul buku yang ada di posisi teratas langsung dikeluarin hasilnya ke layar. Terakhir, kita input satu angka rating buku yang mau dicari, dan program bakal nyari posisi bukunya pakai metode Binary Search alias pencarian belah dua biar nyarinya cepet. Kalau ketemu, atribut bukunya langsung dikeluarin datanya pakai Printf, tapi kalau zonk, dia bakal ngeprint tulisan "Tidak ada buku dengan rating seperti itu"

