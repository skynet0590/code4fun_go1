package main

import (
	"fmt"
	"time"
)

// https://tour.golang.org/
// https://godoc.org/golang.org/x/net/html
// https://github.com/asdine/storm#getting-started
// https://www.cockroachlabs.com/

func main() {
	go func() {

	}()
	var num int
	for {
		num ++
		if (num % 2) == 0 {
			continue
		}
		fmt.Printf("For %v time(s) \n", num)
		if num == 5 {
			break
		}
		time.Sleep(time.Second * 2)
	}
}
