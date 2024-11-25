package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/// escrita
	f, err := os.Create("arquivo.txt")

	// tamanho, err := f.WriteString("Hello, World!")
	tamanho, err := f.Write([]byte("Escrevendo dados no arquivo"))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Arquivo criado com sucesso! Tamanho: %d bytes\n", tamanho)
	f.Close()

	/// leitura

	// arquivo, err := os.Open(("arquivo.txt"))
	arquivo, err := os.ReadFile(("arquivo.txt"))

	if err != nil {
		panic(err)
	}

	fmt.Printf(string(arquivo))

	// leitura de pouco em pouco abrindo o arquivo

	arquivo2, err2 := os.Open("arquivo.txt")

	if err2 != nil {
		panic((err2))
	}

	reader := bufio.NewReader(arquivo2)
	buffer := make([]byte, 10)

	for {
		n, err := reader.Read(buffer)

		if err != nil {
			break
		}

		fmt.Println(string(buffer[:n]))
	}

	err = os.Remove("arquivo.txt")

	if err != nil {
		panic(err)
	}
}
