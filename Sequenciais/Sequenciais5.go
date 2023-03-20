package main

import "fmt"

func main() {
	var idade int
	fmt.Println("digite sua idade ")
	fmt.Scan(&idade)
	fmt.Println("Sua idade em dias é", idade*365)
}
