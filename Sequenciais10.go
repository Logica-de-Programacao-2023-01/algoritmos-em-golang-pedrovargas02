package main

import "fmt"

func main() {
	var peso float64
	fmt.Println("Informe um peso, em kg, para ser convertido em libras")
	fmt.Scan(&peso)
	fmt.Println("O peso em libras é", peso*2.205, "lb")
}
