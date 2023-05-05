package main

import "fmt"

func main() {
	i := 5

	switch i {
	case 1:
		fmt.Println("i = 1")

	case 2:
		fmt.Println("i = 2")

	case 5:
		fmt.Println("i = 5")

	default:
		fmt.Println("i değeri bilinmiyor.")
	}

	// Koşulsuz switch
	x := 10

	switch {
	case x == 10:
		fmt.Println("x = 10") // ilk koşul doğru olduğu için diğer koşullara bakılmaz.
	case x < 20:
		fmt.Println("x < 20")
	}

	// bir alttaki koşulun kontrol edilmesi için:
	y := 10
	switch {
	case y == 10:
		fmt.Println("y = 10")
		fallthrough
	case y < 20:
		fmt.Println("y < 20")

	}
}
