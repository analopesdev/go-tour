package main

//interfaces

type Pessoa interface {
	Desativar()
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

// client tem o metodo Desativar implementado
func (c *Cliente) Desativar() {
	c.Ativo = false
}

// qualquer struct que tiver o metodo Desativar implementado, pode ser passado como parametro
func Desativacao(p Pessoa) {
	p.Desativar()
}

func main() {

	// john is people because implements the method Desativar

	cliente := Cliente{
		Nome:  "John",
		Idade: 20,
		Ativo: true,
	}

	Desativacao(&cliente)
}
