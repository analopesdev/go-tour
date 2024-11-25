package main

import "fmt"

const a = "ana"

var u = "teste"

type Id int

var (
	b bool    = true
	c int     = 10
	d string  = "Ana"
	e float64 = 12.0
	f Id      = 1
)

var ana Id = 20

func main() {
	fmt.Printf("O tipo de é %T", f)
}
