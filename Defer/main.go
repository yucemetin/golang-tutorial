package main

import "fmt"

// defer kullanılan satır, içinde bulunduğu fonksiyonda en son işleme alınır.
func main() {
	defer fmt.Println("\nİlk cümle")
	fmt.Println("İkinci cümle")

	s := "Metin"
	s1 := []rune(s)

	for _, n := range s1 {
		defer fmt.Print(string(n))
	}
}
