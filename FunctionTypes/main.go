package main

import "fmt"

// Variadic Functions
func add(nums ...int) int {
	sum := 0

	for _, n := range nums {
		sum += n
	}

	return sum
}

// fonksiyon, baska bir fonksiyon döndürebilir.
func printS() func(s string) {
	return func(s string) {
		fmt.Println(s)
	}
}

// Recursive Function
func fac(num int) int {
	if num == 0 {
		return 1
	}

	return num * fac(num-1)
}

// 2 Returns Functions
func division(num1, num2 int) (int, int) {
	p1 := num1 / 2
	p2 := num2 / 2

	return p1, p2
}

func main() {
	fmt.Println(add(1, 2, 3, 4, 5, 6))

	// Closure Functions
	add2 := func(num1, num2 int) int {
		return num1 + num2
	}

	fmt.Println(add2(5, 15))

	x := printS()
	x("Deneme")

	fmt.Println(fac(5))

	// Anonymous Function

	s := "Metin"

	func(a string) {
		fmt.Println(a)
	}(s)

	w, q := division(10, 20)
	fmt.Println(w, q)

	// Eğer değerlerden birisini almak istemiyorsak "_" kullanılır
	num1, _ := division(100, 20)

	fmt.Println(num1)
}
