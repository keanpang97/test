package main

import (
	"fmt"
	"functioning"
)

func main() {
	greeting := functioning.Hello("Kean Pang")
	fmt.Println(greeting)

	total := functioning.Sum(4, 5)
	fmt.Println(total)
}
