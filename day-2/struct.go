package main

import (
	"fmt"
)

type (
	User struct {
		Id int
		Name string
		Password string
	}
	MyNum int
)

func (u User) GetName() string {
	return "Hello My name is: " + u.Name
}

func (i *MyNum) Plus(in int) int {
	*i += MyNum(in)
	fmt.Println(i)
	return int(*i)
}

func main() {
	/*dog := child.Dog{
		Id:   0,
		Name: "",
	}
	fmt.Printf("%+v \n", dog)*/
	num := MyNum(6)
	fmt.Println(num)
	fmt.Println(num.Plus(5))
	fmt.Println(num)

	num2 := &num
	fmt.Println(num2)
	// cai nay chay dc
	/*cat := child.cat{
		Id:   0,
		Name: "",
	}
	fmt.Printf("%+v \n", cat)*/
	// cai nay se loi
}
