# <h1 align="center">Laporan Praktikum Modul 4 - ... </h1>
<p align="center">[Muhammad Hudzaifah] - [109082530017]</p>

## SOAL

### 1. [Soal 1]
#### soal1.go

```go
package main

import "fmt"

func factorial(n int, hasil *int) {
	*hasil = 1
	for i := 1; i <= n; i++ {
		*hasil *= i
	}
}

func permutation(n, r int, hasil *int) {
	var nFact, nrFact int
	factorial(n, &nFact)
	factorial(n-r, &nrFact)
	*hasil = nFact / nrFact
}

func combination(n, r int, hasil *int) {
	var nFact, rFact, nrFact int
	factorial(n, &nFact)
	factorial(r, &rFact)
	factorial(n-r, &nrFact)
	*hasil = nFact / (rFact * nrFact)
}

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	var p1, c1, p2, c2 int

	permutation(a, c, &p1)
	combination(a, c, &c1)

	permutation(b, d, &p2)
	combination(b, d, &c2)

	fmt.Printf("%d %d\n", p1, c1)
	fmt.Printf("%d %d\n", p2, c2)
}
```
### Output :

##### Output 
![Screenshot Output soal 1](https://github.com/Halzionns/109082530017_Muhammad-Hudzaifah/blob/main/modul4/output/output1.png)
[penjelasan] 
Pertama, kita siapin dulu fungsi factorial, permutation, dan combination pakai parameter pointer alias *int biar bisa niruin prosedur in/out persis kayak yang diminta di soal, kita bikin logika faktorialnya pakai perulangan biasa buat ngaliin angka berurutan, terus di fungsi permutasi dan kombinasi kita tinggal panggil aja fungsi faktorial tadi buat nerapin rumus pembagiannya, habis itu di fungsi utaman kita scan 4 input angka asli berurutan buat dapet nilai a, b, c, dan d, terakhir, nilai-nilai itu tinggal kita lempar ke dalam prosedur buat dicari hasilnya lalu diprint jadi dua baris buat nampilin hasil dari a terhadap c dan b terhadap d

## SOAL

### 2. [Soal 2]
#### soal2.go

```go
package main

import "fmt"

func hitungSkor(soal *int, skor *int) {
	*soal = 0
	*skor = 0
	var waktu int
	
	for i := 0; i < 8; i++ {
		fmt.Scan(&waktu)
		if waktu < 301 {
			*soal++
			*skor += waktu
		}
	}
}

func main() {
	var nama string
	var pemenangNama string
	pemenangSoal := -1
	pemenangSkor := 999999 

	for {
		fmt.Scan(&nama)
		
		if nama == "Selesai" {
			break
		}

		var soal, skor int
		hitungSkor(&soal, &skor)

		if soal > pemenangSoal {
			pemenangNama = nama
			pemenangSoal = soal
			pemenangSkor = skor
		} else if soal == pemenangSoal {
			if skor < pemenangSkor {
				pemenangNama = nama
				pemenangSoal = soal
				pemenangSkor = skor
			}
		}
	}
	fmt.Printf("%s %d %d\n", pemenangNama, pemenangSoal, pemenangSkor)
}
```
### Output :

##### Output 
![Screenshot Output Soal 2](https://github.com/Halzionns/109082530017_Muhammad-Hudzaifah/blob/main/modul4/output/output2.png)
[penjelasan] 
Pertama, kita siapin dulu prosedur hitungSkor buat ngitung jumlah soal sama total skornya sekaligus ngescan 8 input waktu dari tiap peserta yang mana waktu di bawah 301 menit aja yang dihitung, terus di program utama kita bikin perulangan buat ngebaca nama peserta satu per satu, kalau namanya kebaca Selesai, program langsung nge break perulangannya, nah tiap dapet nama peserta kita panggil prosedur hitungSkor tadi buat dapetin hasilnya, habis itu kita adu nilainya sama rekor pemenang sementara dengan aturan yang paling banyak ngerjain soal atau yang waktunya paling dikit kalau jumlah soalnya seri bakal otomatis jadi kandidat juara baru, terakhir, data juara akhirnya berupa nama peserta beserta jumlah soal dan total waktunya tinggal diprint satu baris di akhir

