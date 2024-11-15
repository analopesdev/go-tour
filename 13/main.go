package main

import "fmt"

type Cliente struct {
	nome string
}

func (c *Cliente) andou() {
	c.nome = "Ana Lopes"
	fmt.Printf("O cliente %v andou", c.nome)
}

func main() {

	client1 := Cliente{
		nome: "Ana",
	}

	client1.andou()
	fmt.Printf("O valor da stuct com o nome %v", client1.nome)
}
