package main

// struct metods
type Client struct {
	Name string
	Age  int
}

type People struct {
	Name string
	Age  int
}

type ControllerTest struct {
	service ServiceTest
}

type ServiceTest struct {
	repo RepositoryInterface
}

type RepositoryInterface interface {
	SetName(name string)
}

type RepositoryTest struct {
	repo RepositoryInterface
}

func (p *People) SetName(name string) {
	p.Name = name
}

func (c *Client) SetName(name string) {
	c.Name = name
}

func (c *ControllerTest) Test() {
	c.service.repo.SetName("Ana")
}

func main() {
	client := Client{
		Name: "John",
		Age:  20,
	}

	client.SetName("Doe")

	people := People{}

	people.SetName("Ana")

	println(client.Name)
}

// para chamar metodos de uma struct, é necessário passar o ponteiro da struct, ex:

// func (p *People) SetName(name string) {
// 	p.Name = name
// }
