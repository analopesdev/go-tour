package main

func main() {

	for i := 0; i < 10; i++ {
		println(i)
	}

	numeros := []string{"um", "dois", "tres"}

	for i, v := range numeros {
		println(v, i)
	}

	i2 := 0

	for i2 < 10 {
		println(i2)
		i2++
	}

	for {
		println("Ola")
	}
}
