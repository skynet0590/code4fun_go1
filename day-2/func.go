package main

import "fmt"

func main() {
	greet, err := hello("World", []int{2, 1, 3}...)
	fmt.Println(greet, err)
}

func hello(name string, nums ...int) (greet string, err error) {
	num := 0
	for _,n := range nums {
		//fmt.Println(i)
		num += n
	}
	greet = fmt.Sprintf("Hello %s %v time(s)", name, num)
	return
}
