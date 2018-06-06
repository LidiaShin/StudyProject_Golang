package main

import "fmt"

func main() {
	fmt.Println(true && true)
	fmt.Println("true and false : ", true && false)
	fmt.Println("false and false : ", false && false)

	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)

	fmt.Println("321325 × 424521 : ", 321325*424521)
}
