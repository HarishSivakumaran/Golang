package main

import (
	"fmt"
)

// walrus operator can't be written outside func block
var PublicVar string = "qwertyyyyyyy"
var privVar string = "qwertyyyyyyy"

func main() {
	var num uint16 = 1698
	var name string = "Harish"
	weight := 77.77

	fmt.Println(num)
	fmt.Println(name)
	fmt.Println(weight)
	fmt.Println(PublicVar)
	fmt.Println(privVar)
}
