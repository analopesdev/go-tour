package main

func main() {
	a := 5
	b := 2

	if a < b {
		println(a)

	} else {
		println(b)
	}

	switch a {
	case 1:
		println("a")
	case 2:
		println("a")

	default:
		println("b")
	}

	if a > b && b < a {
		println("a maior que b")
	}
}
