package main

import (
	"fmt"
)

func main() {
	var peso float64
	var altura float64
	fmt.Println("Vamos calcular o seu IMC")
	fmt.Println("Informe o seu peso em kg")
	fmt.Scan(&peso)
	if peso <= 0 {
		fmt.Println("Informe um valor maior que 0")
		return
	}
	fmt.Println("Informe sua altura em metros")
	fmt.Scan(&altura)
	if altura <= 0 {
		fmt.Println("Informe um valor maior que 0")
		return
	}
	fmt.Println("O seu IMC é", peso/(altura*altura))
}
