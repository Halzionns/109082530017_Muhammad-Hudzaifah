# <h1 align="center">Laporan Praktikum Modul 3 - ... </h1>
<p align="center">[Muhammad Hudzaifah] - [109082530017]</p>

## Unguided 

### 1. [Soal 1]
#### soal1.go

```go
package main

import "fmt"

package main

import "fmt"

func factorial(n int) int {
	res := 1
	for i := 2; i <= n; i++ {
		res *= i
	}
	return res
}

func permutation(n, r int) int {
	return factorial(n) / factorial(n-r)
}

func combination(n, r int) int {
	return factorial(n) / (factorial(r) * factorial(n-r))
}

func main() {
	var a, b, c, d int

	_, err := fmt.Scan(&a, &b, &c, &d)
	if err != nil {
		return
	}

	fmt.Printf("%d %d\n", permutation(a, c), combination(a, c))

	fmt.Printf("%d %d\n", permutation(b, d), combination(b, d))
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Halzionns/109082530017_Muhammad-Hudzaifah/modul3/output/output-soal1.png)
[penjelasan]
Kombinasi dan permutasi

## Unguided 

### 2. [Soal 2]
#### soal2.go

```go
package main

import "fmt"

func f(x int) int {
    return x * x 
}

func g(x int) int {
    return x - 2 
}

func h(x int) int {
    return x + 1 
}

func main() {
    var a, b, c int
    
    fmt.Scan(&a, &b, &c)

    fmt.Println(f(g(h(a))))

    fmt.Println(g(h(f(b))))

    fmt.Println(h(f(g(c))))
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Halzionns/109082530017_Muhammad-Hudzaifah/modul3/output/output-soal2.png)
[penjelasan]
penulisan dalam bentuk function

## Unguided 

### 3. [Soal 3]
#### soal3.go

```go
package main

import (
	"fmt"
	"math"
)

func jarak(a, b, c, d float64) float64 {
	return math.Sqrt(math.Pow(a-c, 2) + math.Pow(b-d, 2))
}

func didalam(cx, cy, r, x, y float64) bool {
	return jarak(cx, cy, x, y) <= r
}

func main() {
	var cx1, cy1, r1 float64
	var cx2, cy2, r2 float64
	var x, y float64

	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Scan(&cx2, &cy2, &r2)
	fmt.Scan(&x, &y)

	diLingkaran1 := didalam(cx1, cy1, r1, x, y)
	diLingkaran2 := didalam(cx2, cy2, r2, x, y)

	if diLingkaran1 && diLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if diLingkaran1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if diLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Halzionns/109082530017_Muhammad-Hudzaifah/modul3/output/output-soal3.png)
[penjelasan]
koordinat