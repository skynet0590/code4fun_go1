package main

import "fmt"

func main() {
	var p *int
	a := 8
	p = &a
	a = 10
	fmt.Println(*p)
}