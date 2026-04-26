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