package main

import "fmt"

func Soma(a int, b int) int {
	return a + b
}

func main() {

	var a int = 5
	var b int = 5

	total := Soma(a, b)

	fmt.Printf("A soma de %d + %d é igual %d\n", a, b, total)
}
