# <h1 align="center">Laporan Praktikum Modul 10 - ... </h1>
<p align="center">[Muhammad Hudzaifah] - [109082530017]</p>

## SOAL

### 1. [Soal 1]
#### soal1.go

```go
package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Masukkan jumlah kelinci : ")
	fmt.Scan(&n)

	var beratKelinci [1000]float64

	fmt.Print("Masukkan berat beratnya : ")
	
	for i := 0; i < n; i++ {
		fmt.Scan(&beratKelinci[i])
	}

	if n > 0 {
		min := beratKelinci[0]
		max := beratKelinci[0]

		for i := 1; i < n; i++ {
			if beratKelinci[i] < min {
				min = beratKelinci[i]
			}
			if beratKelinci[i] > max {
				max = beratKelinci[i]
			}
		}
		fmt.Printf("\nTeringan : %v\n", min)
		fmt.Printf("Terberat : %v\n", max)
	}
}
```
### Output :

##### Output 
![Screenshot Output soal 1](https://github.com/Halzionns/109082530017_Muhammad-Hudzaifah/blob/main/modul10/output/output1.png)
[penjelasan] Pertama tama, pada scan n, kita menginput jumlah kelinci yang akan kita masukkan beratnya, pada beratKelinci[1000]float64, kita membuat Array dengan kapasitas seribu, bjir.... Didalam loopnya, kita akan menginput berat berat kelinci sesuai dengan angka yang kita masukkan pada input sebelumnya, dan disitu, input pertama akan dianggap sebagai yang terberat dan teringan untukl sementara, dan loopnya akan mengecheck apakah ada angka yang lebih kecil atau lebih besar, ketika ketemu, maka yang terkecil akan ditetapkan, yang terbesar juga akan ditetapkan. Pada Printf nya, kita akan menghasilkan output berdasarkasn data yang terinput dan sudah terolah di perulangan tersebut (\n untuk ganti baris, %v untuk menampung variabel min dan max)


## SOAL

### 2. [Soal 2]
#### soal2.go

```go
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
```
### Output :

##### Output 
![Screenshot Output Soal 2](https://github.com/Halzionns/109082530017_Muhammad-Hudzaifah/blob/main/modul10/output/output2.png)
[penjelasan] Pertama, kita akan input jumlah ikan sama kapasitas wadahnya di Scan &y &x, terus ngebuat Array beratIkan[1000] buat nampung berat-berat ikannya, kapasitasnya jujur banyak banget, kenapa harus seribu yak... Terus di dalam looping nya, kita masukin berat ikannya satu-satu dan langsung dikelompokkin ke wadah pake modulob jadi tiap kali ikannya udah penuh sesuai kapasitas wadah atau emang ikannya udah abis, program bakal langsung ngeprint total berat wadah itu terus timbangannya dibikin nol lagi buat wadah berikutnya, abis itu hasilnya bakal di keluarin pake Printf, baris pertama buat total berat tiap wadah, dan baris kedua buat rata-ratanya hasil bagi total berat semua ikan dengan jumlah wadah yang kepake, terus pake %.2f biar angka desimalnya cuma dua di belakang koma


## SOAL

### 3. [Soal 3]
#### soal3.go

```go
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
```
### Output :

##### Output 
![Screenshot Output Soal 3](https://github.com/Halzionns/109082530017_Muhammad-Hudzaifah/blob/main/modul10/output/output3.png)
[penjelasan] Pertama, kita input ada berapa banyak balita yang mau didata pakai Scan(&n), terus kita siapin array arrBalita buat nampung semua beratnya. Terus dalem looping, kita masukin berat balitanya satu-satu ke dalam array tadi. Biar kita tahu siapa yang paling kecil dan siapa yang paling gede, kita pakai prosedur MinMax. Abis itu kita panggil fungsi bjir buat jumlahin semua berat yang ada terus dibagi sama total balitanya biar dapet nilai rata-ratanya. Terakhir, semua hasilnya kita keluarin pakai Printf dan pake %.2f biar angka desimal di belakang komanya cuma dua
