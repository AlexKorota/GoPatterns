package main

import "fmt"

type (
	AFCreator interface {
		Create(input string) AFProduct
		Register(product AFProduct)
	}

	AFProduct interface {
		Use()
	}

	AFRealCreator struct {
		products []AFProduct
	}

	AFProductA struct {
		command string
	}
	AFProductB struct {
		command string
	}
	AFProductC struct {
		command string
	}
)

func (rc *AFRealCreator) Create(input string) AFProduct {
	var Prod AFProduct
	switch input {
	case "A":
		Prod = AFProductA{input}
	case "B":
		Prod = AFProductB{input}
	case "C":
		Prod = AFProductC{input}
	}
	rc.Register(Prod)
	return Prod
}

func (rc *AFRealCreator) Register(prod AFProduct) {
	rc.products = append(rc.products, prod)
}

func (a AFProductA) Use() {
	fmt.Println(a.command)
}

func (b AFProductB) Use() {
	fmt.Println(b.command)
}

func (c AFProductC) Use() {
	fmt.Println(c.command)
}

func main() {
	cr := new(AFRealCreator)
	cr.Create("A")
	cr.Create("B")
	qq := cr.Create("C")
	cr.Create("A")
	c := cr.products[1]
	c.Use()
	qq.Use()
	fmt.Println(cr)
}
