package main

import (
	"fmt"
	"github.com/skynet0590/code4fun_go1/day-2/child"
	"os"
)

func main() {
	a, err := child.NewAnimal("dog", "Tom")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(a.Said())
	fmt.Println(a.Said())
	fmt.Printf("The annimal have said: %+v time(s) \n", a.SaidNumber())

	c, err := child.NewAnimal("cat", "Terry")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(c.Said())
	fmt.Printf("The annimal have said: %+v time(s) \n", c.SaidNumber())

	dog := child.Dog{
		Id:   0,
		Name: "Ngao Ngo",
	}
	fmt.Println(dog.Said())
	fmt.Printf("The annimal have said: %+v time(s) \n", dog.SaidNumber())

	tiger, err := child.NewAnimal("tiger", "Ngao")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(tiger.Said())
	fmt.Printf("The annimal have said: %+v time(s) \n", tiger.SaidNumber())


}