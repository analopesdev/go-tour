package main

// func main() {
// 	// Memoria -> endereço -> Valor

// 	a := 10
// 	println(&a)

// 	// Atribuindo o endereco da memoria de a em b
// 	var ponteiro *int = &a
// 	// mudar valor de ponteiro
// 	*ponteiro = 20
// 	println(a) // 20

// 	// apontar para o endereco de a
// 	b := &a

// 	println(b) //0xc000066738

// 	// * aponta para o valor
// 	println(*b) // 20

// 	*b = 30

// 	println(*b)
// 	println(a)
// }

func main() {
	// memoria -> endereco -> valor

	a := 10

	// print endereco de a
	println(&a)

	var b *int = &a
	*b = 30

	// c agora aponta para o mesmo endereco de b
	c := &b

	println(*c)
	println(b)
	// * representa valor
	println(*b)
}
