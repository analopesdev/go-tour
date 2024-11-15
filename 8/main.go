package main

import (
	"fmt"
)

func main() {

	fmt.Println(sum(1, 2, 4, 6))

	total := func() int {
		return 2
	}()

	fmt.Println(total)
}

func sum(numeros ...int) int {
	total := 0

	for _, num := range numeros {
		total += num
	}

	return total
}
