package main

import "fmt"

func add(num1 int, num2 int) int {
	return num1 + num2
}

func add2(num1, num2 int) int { // eğer parametredeki değerlerin hepsi aynı tür ise tek bir int anahtarı yeterlidir.
	return num1 + num2
}

func function(num int) (x, y int) { // parantez içinde return edilecek değişkenler tanımlanabilir.
	x = num / 2
	y = num * 2
	return // sadece return kullanılırsa tanımlanmış olan değişkenler otomatik olarak return edilir.
}

func function2(num int) (x, y, z int) {
	x = num
	y = num * 2
	z = num * 3
	return
}

func main() {
	var num1 = 36
	var num2 = 14
	const num3 = 50
	num4 := 20

	fmt.Println(add(num1, num2))
	fmt.Println(add2(num3, num4))

	fmt.Println(function(10))
	fmt.Println(function2(10))

}
