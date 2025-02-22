package main

func main() {
	// funcoes variaticas
	println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	println(sum(1, 2, 3, 4, 5))
	println(sum(1, 2, 3, 4))
	println(sum(1, 2, 3))
	println(sum(1, 2))
	println(sum(1))
	println(sum())
}

// funcoes variaticas
func sum(numeros ...int) int {
	var total int

	for _, numero := range numeros {
		total += numero
	}

	return total
}

// closures
func closure() {
	var x int
	x = 10

	//anonimous function
	func() {
		println(x)
	}()
}
