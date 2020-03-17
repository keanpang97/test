package main

import (
	"fmt"
	"function"
)

func main() {
	greeting := function.Hello("Kean Pang")
	fmt.Println(greeting)

	total := function.Sum(4, 5)
	fmt.Println(total)
}
