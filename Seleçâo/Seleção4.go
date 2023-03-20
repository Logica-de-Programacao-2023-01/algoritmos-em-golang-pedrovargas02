package main

import "fmt"

func main() {
	var peso float64
	var altura float64
	var sexo float64
	var IMC float64

	fmt.Print("Informe, respectivamente,seu peso (em kg), sua altura(em metros) e sexo ")
	fmt.Scan(&peso, &altura, &sexo)

	IMC = peso / (altura * altura)

	if IMC < 18.5 {
		fmt.Println("Você está abaixo do peso ideal")
	} else if 18.5 <= IMC && IMC <= 24.99 {
		fmt.Println("Você está dentro do peso ideal")
	} else {
		fmt.Println("Você está acima do peso ideal")
	}
}
