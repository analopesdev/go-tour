package main

const a = "Hello, word!"

// declaracao de escopo global
var (
	b bool
	c string
	d int
	e float64
)

func main() {
	b = true
	c = "Hello, word!"
	d = 10
	e = 10.5

	//short hand

	println(b)
}
