package main

import "fmt"

func main() {
	var valor float64
	fmt.Println("Informe o valor do produto")
	fmt.Scan(&valor)
	fmt.Println("Com um desconto de 10% o valor do seu produto será", valor-valor*10/100)
}
