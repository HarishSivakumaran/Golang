package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("write anything")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("you wrote,", input)
}
