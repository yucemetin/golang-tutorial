package main

import "fmt"

func main() {

	var names = []string{"metin", "ayse", "mehmet", "mustafa"}

	for i := 0; i < 4; i++ {
		fmt.Println(names[i])
	}

	for _, n := range names {
		fmt.Println(n)
	}

	// ilk sayi indexi ikinci ise değeri temsil eder.
	for i, n := range names {
		fmt.Printf("%d. index: %s\n", i, n)
	}

}
