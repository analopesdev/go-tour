package main

func main() {
	var minhaVar interface{} = "Ana"

	println(minhaVar.(string))

	res, ok := minhaVar.(int)
	res2 := minhaVar.(int)
	println(res2)
	println(res)
	println(ok)
}
