package main

import "fmt"

type (
	Creator interface {
		Create(input string) Product
		Register(product Product)
	}

	Product interface {
		Use()
	}

	RealCreator struct {
		products []Product
	}

	ProductA struct {
		command string
	}
	ProductB struct {
		command string
	}
	ProductC struct {
		command string
	}
)

func (rc *RealCreator) Create(input string) Product {
	var Prod Product
	switch input {
	case "A":
		Prod = ProductA{input}
	case "B":
		Prod = ProductB{input}
	case "C":
		Prod = ProductC{input}
	}
	rc.Register(Prod)
	return Prod
}

func (rc *RealCreator) Register(prod Product) {
	rc.products = append(rc.products, prod)
}

func (a ProductA) Use() {
	fmt.Println(a.command)
}

func (b ProductB) Use() {
	fmt.Println(b.command)
}

func (c ProductC) Use() {
	fmt.Println(c.command)
}

func main() {
	cr := new(RealCreator)
	cr.Create("A")
	cr.Create("B")
	qq := cr.Create("C")
	cr.Create("A")
	c := cr.products[1]
	c.Use()
	qq.Use()
	fmt.Println(cr)
}
