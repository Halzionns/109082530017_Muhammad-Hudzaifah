# <h1 align="center">Laporan Praktikum Modul 3 - ... </h1>
<p align="center">[Muhammad Hudzaifah] - [109082530017]</p>

## Unguided 

### 1. [Soal]
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
![Screenshot Output Unguided 1_1](https://github.com/Halzionns/109082530017_Muhammad-Hudzaifah/Modul3/output/Output1.pbn)
[penjelasan]

