package child

import (
	"errors"
	"fmt"
)

func init() {
	fmt.Println("child was inited")
}

type (
	Dog struct {
		Id int
		saidNumber int
		Name string
	}
	cat struct {
		Id int
		saidNumber int
		Name string
	}
	Animal interface {
		Said() string
		SaidNumber() int
	}
)


func (c *cat) SaidNumber() int {
	return c.saidNumber
}

func (d *Dog) SaidNumber() int {
	return d.saidNumber
}

func (c *cat) Said() string {
	c.saidNumber ++
	return fmt.Sprintf("%s said: Mew mew", c.Name)
}

func (d *Dog) Said() string {
	d.saidNumber ++
	return fmt.Sprintf("%s said: Gau Gau", d.Name)
}

func NewAnimal(animalType string, name string) (a Animal, err error) {
	if animalType == "dog" {
		return &Dog{
			Id:   0,
			Name: name,
		}, nil
	}
	if animalType == "cat" {
		return &cat{
			Id:   0,
			Name: name,
		}, nil
	}
	return nil, errors.New("Unsupportd animal type")
}