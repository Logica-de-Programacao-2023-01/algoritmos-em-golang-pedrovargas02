package main

import "fmt"

func main() {
	var x int
	fmt.Print("Informe um número ")
	fmt.Scan(&x)

	if x%2 == 0 {
		fmt.Println("O número é par")
	} else if x%2 != 0 {
		fmt.Println("O número é impar")
	}
}
