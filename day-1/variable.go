package main

import "fmt"

func main() {
	num := 0
	var f = 0.0
	info(num)
	info(f)
	// info(int8(5555555555555555555)) // Cai nay se loi
	// info(int32(5555555555555555555)) // Cai nay se loi
	info(int64(5555555555555555555))
	info(5555555555555555555) // Cai nay se loi neu ban chay OS x84 (32 bit)
	info("")
	info(true)
	info(0xf3) // hexa
	info(1<<9)
	info(1<<8)
	info(3<<8) // = 3 * 1<<8
	info(1e9)
	info(1e3 == float64(1000))
	info(make(map[int]string))
	var m map[int]string
	info(m)
	info(map[int]string{
		1: "So 1",
		2: "Hello",
	})
	info([]int{0,1,2})
}

func info(v interface{}) {
	fmt.Printf("Type: %T; Value: %+v \n", v, v)
}
