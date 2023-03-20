package main

import "fmt"

func main() {
	var numero1 float64
	var numero2 float64
	var numero3 float64
	fmt.Println("Informe 3 números para ser mostrada a média ponderada entre eles, com peso 2, 3 e 5, respectivamente ")
	fmt.Scan(&numero1, &numero2, &numero3)
	fmt.Println("Aqui está:", (numero1*2+numero2*3+numero3*5)/(numero1+numero2+numero3))
}
