package main

import "fmt"

func main() {

	test("Ana")                     // []string
	test("Ana", "Maria", "Beatriz") // []string
	//parametros veridiacos
}

func test(name string, sobrenome ...string) {
	fmt.Println(name)
	fmt.Println(sobrenome) // []string
}
