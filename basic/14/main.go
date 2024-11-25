package main

import "fmt"

// implements all
// type x interface{}

func main() {
	var x interface{} = 10
	var y interface{} = "Ana"

	showType(x)
	showType(y)
}

func showType(t interface{}) {
	fmt.Printf("O tipo da variavel e %T e o valor e %v\n", t, t)
}
