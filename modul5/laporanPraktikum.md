# <h1 align="center">Laporan Praktikum Modul 5 - ... </h1>
<p align="center">[Muhammad Hudzaifah] - [109082530017]</p>

## SOAL

### 1. [Soal 1]
#### soal1.go

```go
package main

import "fmt"

func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	var n int
	
	fmt.Print("Masukkan nilai n: ")
	fmt.Scan(&n)
	
	hasil := fibonacci(n)
	fmt.Printf("Nilai deret Fibonacci ke-%d adalah %d\n", n, hasil)
}
```
### Output :

##### Output 
![Screenshot Output soal 1](https://github.com/Halzionns/109082530017_Muhammad-Hudzaifah/blob/main/modul5/output/output1.png)
[penjelasan] Pertama, kita siapin dulu fungsi fibonacci pakai parameter integer buat nampung nilai n, terus kita kasih kondisi dasar atau base case di dalamnya, kalau nilai n sama dengan 0 dia bakal balikin nilai 0, dan kalau n sama dengan 1 dia bakal balikin nilai 1, habis itu buat nilai n yang lebih besar, kita bikin dia memanggil dirinya sendiri alias rekursif buat ngejumlahin hasil dari dua suku sebelumnya secara otomatis, terakhir di fungsi utama kita tinggal scan input angka dari user lalu angkanya dilempar ke fungsi fibonacci tadi dan hasilnya langsung diprint buat nampilin nilai deretnya



## SOAL

### 2. [Soal 2]
#### soal2.go

```go
package main

import "fmt"

func cetakBintang(n int) {
	if n == 0 {
		return
	}
	
	cetakBintang(n - 1)
	
	for i := 0; i < n; i++ {
		fmt.Print("*")
	}
	fmt.Println()
}

func main() {
	var n int
	
	fmt.Scan(&n)
	
	cetakBintang(n)
}
```
### Output :

##### Output 
![Screenshot Output Soal 2](https://github.com/Halzionns/109082530017_Muhammad-Hudzaifah/blob/main/modul5/output/output2.png)
[penjelasan] Pertama, kita siapin fungsi rekursif cetakBintang pakai parameter integer buat nampung jumlah barisnya, terus kita kasih kondisi dasar di mana kalau nilainya udah nembus 0 dia bakal langsung berhenti alias di return, habis itu kita panggil fungsi itu sendiri dengan ngurangin nilai n-nya satu per satu biar programnya numpuk instruksi dan mulai nyetak dari baris yang bintangnya paling dikit dulu, nah setelah pemanggilan rekursifnya kita taruh perulangan biasa buat nge print bintang sebanyak nilai n di baris tersebut lalu ditambahin enter biar turun ke bawah, terakhir, di program utama kita tinggal scan input angka dari user dan angkanya langsung dilempar ke fungsi cetakBintang tadi buat nampilin hasil akhirnya



## SOAL

### 3. [Soal 3]
#### soal3.go

```go
package main

import "fmt"

func cetakFaktor(n int, i int) {
	if i > n {
		return
	}
	
	if n%i == 0 {
		fmt.Printf("%d ", i)
	}
	
	cetakFaktor(n, i+1)
}

func main() {
	var n int
	
	fmt.Scan(&n)
	
	cetakFaktor(n, 1)
	fmt.Println()
}
```
### Output :

##### Output 
![Screenshot Output Soal 3](https://github.com/Halzionns/109082530017_Muhammad-Hudzaifah/blob/main/modul5/output/output3.png)
[penjelasan]Pertama, kita siapin dulu fungsi rekursif cetakFaktor yang pakai dua parameter integer yaitu n sebagai angka patokan dan i sebagai angka pembaginya, terus kita kasih kondisi berhenti alias base case di mana kalau nilai i udah ngelewatin n dia bakal langsung berhenti, habis itu kita pakai operasi modulus buat ngecek apakah n habis dibagi i dan kalau nilainya nol langsung kita print aja karena itu pasti faktornya, setelah itu fungsinya memanggil dirinya sendiri dengan ningkatin nilai pembagi i ditambah satu biar dia terus maju ngecek angka berikutnya, terakhir, di fungsi utama kita scan inputnya dari user lalu lempar nilainya ke fungsi cetakFaktor tadi dengan nilai awal i diatur jadi angka 1 biar hasilnya nge print berurutan dari angka paling kecil sampai ujung bilangannya

