package main

// func SomaInteiro(m map[string]int) int {
// 	var soma int

// 	for _, v := range m {
// 		soma += v
// 	}

// 	return soma
// }

// func SomaFloat(m map[string]float64) float64 {
// 	var soma float64

// 	for _, v := range m {
// 		soma += v
// 	}

// 	return soma
// }

// Generics
// func Soma[T int | float64](m map[string]T) T {
// 	var soma T

// 	for _, v := range m {
// 		soma += v
// 	}

// 	return soma
// }

// Generics with type

type Number interface {
	~int | ~float64
}

type MyNumber int

func Soma[T Number](m map[string]T) T {
	var soma T

	for _, v := range m {
		soma += v
	}

	return soma
}

// func Compara[T Number](a T, b T) bool {
// 	if a == b {
// 		return true
// 	}

// 	return false
// }

func Compara[T comparable](a T, b T) bool {
	if a == b {
		return true
	}

	return false
}

func main() {
	m := map[string]int{"Ana": 1000, "Joao": 2000, "Maria": 3000}
	m2 := map[string]int{"Ana": 10.00, "Joao": 2.000, "Maria": 3.000}
	m3 := map[string]MyNumber{"Ana": 1000, "Joao": 2000, "Maria": 3000}

	println(Soma(m))
	println(Soma(m2))
	println(Soma(m3))

	println(Compara(10, 10.0))
}
