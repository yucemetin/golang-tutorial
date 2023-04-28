package main

import "fmt"

func main() {

	var age = 26
	var name = "Metin"

	age2 := 26
	name2 := "Metin"

	fmt.Println("Age: ", age, "\n"+"Name: "+name)
	fmt.Printf("Age: %d \nName: %s", age2, name2)

	const name3 = "Metin3"
	//const name4:="Metin" -> const kullanıldığı zaman ":=" yerine "=" kullanılır.

	fmt.Println("\nName: " + name3)

	variable := "bir"
	{ // kod blokları ayrı olduğu için oluşturulan değişkenler birbirinden bağımsız olur.
		// değişkenlerin adresleri de farklıdır.
		variable := "iki"
		fmt.Println(variable)
		fmt.Println(&variable)
	}
	fmt.Println(variable)
	fmt.Println(&variable)

	variable2 := 8378
	variable3 := string(8378)

	fmt.Println(variable2)
	fmt.Println(variable3) // integer bir değişkeni string olarak atarsak karakter olarak deper alır. (₺)

}
