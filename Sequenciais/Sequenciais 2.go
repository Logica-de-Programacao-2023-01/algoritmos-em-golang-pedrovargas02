package main

import "fmt"

func main() {
	var x int
	fmt.Println("Informe um número para ser mostrado o seu dobro, triplo e quadraplo")
	fmt.Scan(&x)
	fmt.Println("Aqui está", x*2, x*3, x*4)
}
