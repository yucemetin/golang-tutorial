package main

import "fmt"

func add(num int) {
	num += 5
}

func addWithPointer(num *int) {
	*num += 5
}

func main() {
	a := 8
	add(a)
	fmt.Println(a) // 8
	addWithPointer(&a)
	fmt.Println(a) // 13
}
