package main

import "fmt"

func main() {
	// a := 4
	// b := 2

	// fmt.Printf("a&b = %v\n", a&b)
	// fmt.Printf("a|b = %v\n", a|b)
	// fmt.Println("a^b =", a^b)

	// a := 21
	// c := a % 10
	// a = a / 10
	// d := a % 10

	// fmt.Printf("첫번째 수 : %v\n두번째 수 : %v\n", c, d)

	a := 4

	fmt.Println(a &^ a)
	fmt.Println(a << 1)
	fmt.Println(a >> 1)

}
