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

func (c Client) Inactive() {
	c.Active = false
	fmt.Printf("O client %s foi desativado", c.Name)
}

func main() {
	ana := Client{
		Name:   "Ana",
		Age:    26,
		Active: true,
	}

	ana.Active = true

	ana.Inactive()
	fmt.Println(ana)
	// // fmt.Println(ana.State)
	// fmt.Println(ana.Address.State)

}
