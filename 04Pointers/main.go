package main

import "fmt"

func main() {
	var ptr *int
	fmt.Println(ptr)
	var num int = 16
	ptr = &num
	fmt.Println(ptr)
	fmt.Println(*ptr)

	*ptr = *ptr * 100
	fmt.Println(ptr)
	fmt.Println(*ptr)

	val := string(12345)

	fmt.Println(val)

	n := int(123)
	n++
}
