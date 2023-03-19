package main

import "fmt"

func main() {
	var numero int
	fmt.Println("informe um numero inteiro")
	fmt.Scan(&numero)
	fmt.Println("o seu sucessor é", numero+1, "e seu antecessor é", numero-1)
}
