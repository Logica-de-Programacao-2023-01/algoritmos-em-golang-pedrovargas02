package main

import "fmt"

func main() {
	var dias float64
	var diaria float64
	fmt.Println("informe o valor da sua diária")
	fmt.Scan(&diaria)
	fmt.Println("Informe os dias que trabalhou")
	fmt.Scan(&dias)
	fmt.Println("O valor do seu salário será", dias*diaria)
}
