package main

import "fmt"

func main() {
	var x int
	var y int
	var z int
	fmt.Println("Informe três números para serem somados")
	fmt.Scan(&x, &y, &z)
	fmt.Println("A soma entre eles será ", x+y+z)

}
