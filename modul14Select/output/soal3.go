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