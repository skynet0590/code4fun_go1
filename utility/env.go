package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("env SESSION_KEY: ", os.Getenv("SESSION_KEY"))
	fmt.Println("env GOROOT: ", os.Getenv("GOROOT"))
}
