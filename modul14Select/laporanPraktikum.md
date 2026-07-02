# <h1 align="center">Laporan Praktikum Modul 14 Selection - ... </h1>
<p align="center">[Muhammad Hudzaifah] - [109082530017]</p>

## SOAL

### 1. [Soal 1]
#### soal1.go

```go
package main

import "fmt"

func selectionSort(arr []int) {
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

func main() {
	var n int
	
	fmt.Scan(&n)
	
	for i := 0; i < n; i++ {
		var m int
		fmt.Scan(&m)
		
		rumah := make([]int, m)
		
		for j := 0; j < m; j++ {
			fmt.Scan(&rumah[j])
		}
		
		selectionSort(rumah)
		
		for j := 0; j < m; j++ {
			fmt.Printf("%d ", rumah[j])
		}
		fmt.Println()
	}
}
```
### Output :

##### Output 
![Screenshot Output soal 1](https://github.com/Halzionns/109082530017_Muhammad-Hudzaifah/blob/main/modul14Select/output/output1.png)
[penjelasan] Pertama, kita bikin fungsi selectionSort dulu buat ngurutin datanya, algoritmanya kerja dengan cara nyari nilai paling kecil di sisa data yang belum terurut, terus dituker posisinya ke depan satu per satu sampai rapi. Terus di program utama, kita scan dulu variabel n buat nentuin berapa banyak daerah yang mau diproses, habis itu kita masuk ke perulangan sebanyak n kali, dan di tiap putarannya kita scan lagi variabel m buat dapetin jumlah rumah, lalu dibikinin slice buat nampung inputan nomor nomor rumahnya. Setelah nomor rumah di satu baris itu udah selesai di scan dan masuk ke slice, tinggal kita lempar datanya ke fungsi selectionSort tadi biar otomatis diurutin membesar, terakhir, slice yang datanya udah rapi tinggal diprint aja pakai perulangan dan spasi, lalu dikasih Println kosong biar output daerah selanjutnya turun ke baris baru


## SOAL

### 2. [Soal 2]
#### soal2.go

```go
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
```
### Output :

##### Output 
![Screenshot Output Soal 2](https://github.com/Halzionns/109082530017_Muhammad-Hudzaifah/blob/main/modul14Select/output/output2.png)
[penjelasan] Pertama, kita siapin dua fungsi sorting yang dimodifikasi. Satu fungsi sortAsc buat ngurutin dari kecil ke besar, dan satu lagi fungsi sortDesc buat ngurutin dari besar ke kecil. Masuk ke program utama, kita scan jumlah daerahnya kayak biasa, terus pas lagi ngescan nomor nomor rumahnya, kita langsung kasih seleksi kondisi, kalau angkanya di modulus 2 hasilnya nggak nol, langsung kita masukin ke slice ganjil, tapi kalau habis dibagi 2 kita masukin ke slice genap, setelah datanya sukses kepisah, kita tinggal panggil fungsinya. Slice ganjil dilempar ke fungsi sortAsc, dan slice genap dilempar ke fungsi sortDesc, kalau udah rapi semua, tinggal kita print pakai perulangan, cetak yang ganjil dulu sampai habis, baru lanjut nyetak yang genap


## SOAL

### 3. [Soal 3]
#### soal3.go

```go
package main

import (
	"fmt"
	"sort"
)

func main() {
	var data []int
	var nom int

	for {
		fmt.Scan(&nom)

		if nom == -5313 {
			break
		}

		if nom == 0 {
			n := len(data)
			if n == 0 {
				continue
			}

			sort.Ints(data)

			if n%2 != 0 {
				fmt.Println(data[n/2])
			} else {
				fmt.Println((data[(n/2)-1] + data[n/2]) / 2)
			}
		} else {
			data = append(data, nom)
		}
	}
}
```
### Output :

##### Output 
![Screenshot Output Soal 3](https://github.com/Halzionns/109082530017_Muhammad-Hudzaifah/blob/main/modul14Select/output/output3.png)
[penjelasan] Pertama, kita siapin slice dinamis buat nampung angkanya, terus kita looping tanpa henti buat ngescan input satu per satu, pas program baca angka -5313, perulangannya langsung kita break biar berhenti total, tapi kalau yang kebaca angka 0, datanya langsung kita urutin pakai fungsi bawaan sort.Ints dan dicek total elemennya, kalau ganjil langsung kita print index tengahnya, kalau genap kita tambahin dua nilai di tengah lalu dibagi dua biar dapet rata-rata yang otomatis dibulatkan ke bawah, kalau angkanya bukan 0 dan bukan -5313, angkanya tinggal kita append terus masukin ke dalam slice penampung tadi buat bahan hitungan selanjutnya
