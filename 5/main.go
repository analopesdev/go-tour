package main

func main() {
	// entendendo ponteiros
	// endereco -> valor na memoria

	a := 10

	println(&a) // endereco de a
	println(a)  // valor de a

	c := a // c recebe o valor de a, c nao é um ponteiro

	// criando um ponteiro aponando para o endereco de a

	b := &a // mesma coisa que b *int = &a

	println(b) // endereco de a
	// b agora é um ponteiro que aponta para o endereco de a

	println(*b) // valor de a

	// mudando o valor de a atraves do ponteiro b

	*b = 20

	println(a)  // valor de a
	println(*b) // valor de a
	println(c)  // valor de a inicio
}
