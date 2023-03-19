package main

import "fmt"

func main() {
	var salario float64
	fmt.Println("Digite o valor do seu sálario")
	fmt.Scan(&salario)
	fmt.Println("Com um aumento de 15% seu salário será", ((salario*15)/100)+salario)
}
