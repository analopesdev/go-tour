package main

// soma recebe uma copia dos valores
func soma(a, b *int) int {
	*a = 50

	return *a + *b
}

func main() {
	//quando usar ponteiros

	minhaVar1 := 10
	minhaVar2 := 20

	// passando copias dos valores
	println(soma(&minhaVar1, &minhaVar2))

	println(minhaVar1)
}
