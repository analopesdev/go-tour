package main

// trabalhar com ponteiros se a ideia é ser multavel
// *int sao valores reais

func soma(a, b *int) int {
	*a = 50
	return *a + *b
}

func main() {

	minhaVar1 := 10
	minhaVar2 := 20

	println(&minhaVar1)

	// passando a copia de minhavar1, minhavar2
	soma(&minhaVar1, &minhaVar2)

	println(minhaVar1)
}
