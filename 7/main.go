package main

import "fmt"

type Cliente struct {
	Nome string
}

type Conta struct {
	saldo int
}

// return address of Conta
func NewConta() *Conta {
	return &Conta{saldo: 0}
}

func (c Cliente) andou() {
	c.Nome = "Ana"
	fmt.Println(c.Nome, "andou")
}

func (c *Conta) simular(valor int) int {
	c.saldo += valor

	return c.saldo
}

func main() {
	cliente := Cliente{
		Nome: "John",
	}

	cliente.andou()

	conta := Conta{
		saldo: 100,
	}

	fmt.Println(cliente.Nome)
	fmt.Println(conta.simular(100))
	fmt.Println(conta.saldo)

	var conta2 *Conta
	conta2 = NewConta()

	fmt.Println(conta2.simular(100))
}
