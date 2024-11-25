package main

import "fmt"

func main() {
	defer fmt.Println("Primeira Linha")
	fmt.Println("Segunda linha")
	fmt.Println("Terceira linha")
}
