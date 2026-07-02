# <h1 align="center">Laporan Praktikum Modul 2 - ... </h1>
<p align="center">[Muhammad Hudzaifah] - [109082530017]</p>

## SOAL

### 1. [Soal 1]
#### soal1.go

```go
package main
import "fmt"
func main() {
var (
satu, dua, tiga string
temp string
)
fmt.Print("Masukan input string: ")
fmt.Scanln(&satu)
fmt.Print("Masukan input string: ")
fmt.Scanln(&dua)
fmt.Print("Masukan input string: ")
fmt.Scanln(&tiga)
fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
temp = satu
satu = dua
dua = tiga
tiga = temp
fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}
```
### Output :

##### Output 
![Screenshot Output soal 1](https://github.com/Halzionns/109082530017_Muhammad-Hudzaifah/blob/main/modul2/output/output1.png)
[penjelasan] 
Pertama, kita siapin 3 variabel string buat data plus 1 variabel temp, kita scan 3 input berurutan, terus print urutan awalnya,
habis itu kita puter posisinya ke kiri, isi var satu dipindah dulu ke temp, terus var dua digeser numpuk ke var satu, dan var tiga digeser numpuk ke dua, terakhir, data awal yang dipindah di temp diselipin masuk ke tiga

## SOAL

### 2. [Soal 2]
#### soal2.go

```go
package main

import "fmt"

func main() {
	var c1, c2, c3, c4 string
	berhasil := true

	for i := 1; i <= 5; i++ {
		fmt.Printf("Percobaan %d: ", i)
		fmt.Scan(&c1, &c2, &c3, &c4)

		if c1 != "merah" || c2 != "kuning" || c3 != "hijau" || c4 != "ungu" {
			berhasil = false
		}
	}

	fmt.Printf("BERHASIL: %t\n", berhasil)
}
```
### Output :

##### Output 
![Screenshot Output Soal 2](https://github.com/Halzionns/109082530017_Muhammad-Hudzaifah/blob/main/modul2/output/output2.png)
[penjelasan] 
Pertama, kita siapin 4 variabel string buat warna dan 1 variabel boolean dengan nilai true, kita looping 5 kali buat scan 4 input berurutan, terus kita cek urutan warnanya tiap percobaan, habis itu kalau ada satu aja susunan yang melenceng dari "merah" "kuning" "hijau" "ungu", nilai booleannya langsung ditimpa jadi false, terakhir, status akhir dari booleannya tinggal diprint buat nampilin hasil true atau false nya

## SOAL

### 3. [Soal 3]
#### soal3.go

```go
package main

import "fmt"

func main() {
	var beratTotal, kg, sisa, biayaKg, biayaSisa, totalBiaya int

	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&beratTotal)

	kg = beratTotal / 1000
	sisa = beratTotal % 1000

	biayaKg = kg * 10000

	if kg >= 10 {
		biayaSisa = 0
	} else {
		if sisa >= 500 {
			biayaSisa = sisa * 5
		} else {
			biayaSisa = sisa * 15
		}
	}

	totalBiaya = biayaKg + biayaSisa

	fmt.Printf("Detail berat: %d kg + %d gr\n", kg, sisa)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaKg, biayaSisa)
	fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)
}
```
### Output :

##### Output 
![Screenshot Output Soal 2](https://github.com/Halzionns/109082530017_Muhammad-Hudzaifah/blob/main/modul2/output/output3.png)
[penjelasan]
Pertama, kita siapin variabel integer buat nampung total berat, berat dalam kg, sisa gram, sama variabel biayanya, kita scan input berat parsel dalam satuan gram, terus kita bagi inputnya dengan 1000 buat dapet nilai kg dan pakai modulus 1000 buat dapet sisa gramnya, habis itu kita hitung biaya dasar dengan cara berat kg dikali 10000, lalu untuk sisa gramnya kita kasih percabangan, kalau berat totalnya lebih dari 10 kg maka biaya sisanya otomatis nol, tapi kalau di bawah itu kita cek lagi sisanya, jika minimal 500 gram dikali 5 rupiah dan jika kurang dari itu dikali 15 rupiah, terakhir, biaya dasar dan sisa dijumlahin terus semuanya tinggal diprint
