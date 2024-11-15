package main

import "fmt"

func main() {
	var meuArray [4]int
	var meuArray3 [10]int

	meuArray[0] = 1
	meuArray[1] = 2
	meuArray[2] = 3
	meuArray[3] = 4

	fmt.Println(len(meuArray))
	fmt.Println(meuArray[len(meuArray)-1])

	for i, v := range meuArray {
		fmt.Printf("O valor do indice é %d é %d\n", i, v)
	}

	for _, v := range meuArray3 {
		fmt.Println(v)
	}
}
