package main

import "fmt"

type Address struct {
	State  string
	Street string
	Number string
}

type Client struct {
	Name   string
	Age    int
	Active bool
	// Address
	Address Address
}

// a
type Pessoa interface {
	// assinatura
	Inactive()
}

type Empresa struct {
	Nome string
}

func (e Empresa) Inactive() {}

// implements client
func (c Client) Inactive() {
	c.Active = false
	fmt.Printf("O client %s foi desativado", c.Name)
}

// qualquer struct que tiver o metodo inactive vai poder ser ultilizado
func onInactive(pessoa Pessoa) {
	pessoa.Inactive()
}

func main() {
	ana := Client{
		Name:   "Ana",
		Age:    26,
		Active: true,
	}

	onInactive(ana)

	ana.Active = true

	ana.Inactive()
	fmt.Println(ana)
	// // fmt.Println(ana.State)
	// fmt.Println(ana.Address.State)

	minhaEmpresa := Empresa{}
	onInactive(minhaEmpresa)

}
