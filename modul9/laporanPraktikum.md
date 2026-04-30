# <h1 align="center">Laporan Praktikum Modul 9 - ... </h1>
<p align="center">[Muhammad Hudzaifah] - [109082530017]</p>

## SOAL

### 1. [Soal 1]
#### soal1.go

```go
package main

import (
	"fmt"
	"math"
)
type Titik struct{ x, y float64 }
type Lingkaran struct {
	pusat Titik
	r     float64
}

func hitungJarak(p1, p2 Titik) float64 {
	return math.Sqrt(math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2))
}

func main() {
	var l1, l2 Lingkaran
	var p Titik

	fmt.Scan(&l1.pusat.x, &l1.pusat.y, &l1.r)
	fmt.Scan(&l2.pusat.x, &l2.pusat.y, &l2.r)
	fmt.Scan(&p.x, &p.y)

	cek1 := hitungJarak(p, l1.pusat) <= l1.r
	cek2 := hitungJarak(p, l2.pusat) <= l2.r

	if cek1 && cek2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if cek1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if cek2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
```
### Output :

##### Output 
![Screenshot Output soal 1](https://github.com/Halzionns/109082530017_Muhammad-Hudzaifah/blob/main/modul9/output/output1.png)
[penjelasan]
Program ini mengecek apakah sebuah titik berada di dalam dua lingkaran tertentu dengan membandingkan jarak titik ke pusat lingkaran terhadap panjang jari-jarinya dengan menggunakan rumus pythagoras, sistem menghitung jarak tersebut, jika jaraknya lebih kecil atau sama dengan jari-jari, maka titik dianggap berada di dalam lingkaran. Hasil akhirnya adalah informasi apakah titik tersebut berada di antara keduanya, hanya masuk di salah satu lingkaran, atau berada di luar keduanya

## SOAL

### 2. [Soal 2]
#### soal2.go

```go
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
```
### Output :

##### Output 
![Screenshot Output Soal 2](https://github.com/Halzionns/109082530017_Muhammad-Hudzaifah/blob/main/modul9/output/output2.png)
[penjelasan]
Program ini membiarkan kita untuk menentukan jumlah indeks dulu sebelum menginput isi Array nya, lalu mennunjukan semua isinya, isi Indeks ganjil, isi Indeks genap, kelipatan, rata rata, deviasi, dan indeks mana yang ingin kita hapus, lalu hasil akhir setelah indeks itu dihapus

## SOAL

### 3. [Soal 3]
#### soal3.go

```go
package main

import "fmt"

func main() {
	var klubA, klubB string
	var skorA, skorB int
	var pemenang []string

	fmt.Print("Klub A : ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B : ")
	fmt.Scan(&klubB)

	i := 1
	for {
		fmt.Printf("Pertandingan %d : ", i)
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			pemenang = append(pemenang, klubA)
		} else if skorB > skorA {
			pemenang = append(pemenang, klubB)
		} else {
			pemenang = append(pemenang, "Draw")
		}
		i++
	}

	fmt.Println("")
	for j, hasil := range pemenang {
		fmt.Printf("Hasil %d : %s\n", j+1, hasil)
	}
	fmt.Println("Pertandingan selesai")
}
```
### Output :

##### Output 
![Screenshot Output Soal 3](https://github.com/Halzionns/109082530017_Muhammad-Hudzaifah/blob/main/modul9/output/output3.png)
[penjelasan]
Program ini membiarkan kita menginput dua nama tim sepakbola, lalu menginput hasil hasil skor tiap pertandingan, lalu program ini akan menunjukkan tim mana yang memenangkabn tiap pertandingannya

## SOAL

### 4. [Soal 4]
#### soal4.go

```go
package main

import "fmt"

func main() {
	var input string
	fmt.Print("Teks: ")
	fmt.Scan(&input)

	karakter := []rune(input)
	n := len(karakter)

	for i := 0; i < n/2; i++ {
		karakter[i], karakter[n-1-i] = karakter[n-1-i], karakter[i]
	}

	teksTerbalik := string(karakter)
	fmt.Println("Reverse teks:", teksTerbalik)

	if input == teksTerbalik {
		fmt.Println("Palindrom: true")
	} else {
		fmt.Println("Palindrom: false")
	}
}
```
### Output :

##### Output 
![Screenshot Output Soal 4](https://github.com/Halzionns/109082530017_Muhammad-Hudzaifah/blob/main/modul9/output/output4.png)
[penjelasan]
Program ini akan mempersilahkan kita untuk menginput sebuah kata lalu membaliknya, dan akan menentukan apakah kata itu palindrom atau bukan