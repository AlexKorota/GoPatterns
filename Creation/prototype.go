package main

import "fmt"

type Prototype interface {
	Clone() Prototype
	GetName() string
}

type Sheep struct {
	name string
}

func (d *Sheep) Clone() Prototype {
	return &Sheep{d.name}
}

func (d *Sheep) GetName() string {
	return d.name
}

func main() {
	sheep1 := Sheep{"Dolly"}
	sheep2 := sheep1.Clone()
	sheep1.name = "Yolly"
	fmt.Println(sheep1.GetName())
	fmt.Println(sheep2.GetName())
}
