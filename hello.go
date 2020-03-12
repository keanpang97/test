package main

import (
	"fmt"
)

func main() {
	greeting := hello("Kean Pang")
	fmt.Println(greeting)

	total := sum(4, 5)
	fmt.Println(total)
}

func hello(user string) string {
	if len(user) == 0 {
		return "Hello Dude!"
	}

	return fmt.Sprintf("Hello %v", user)
}

//addition of two number
func sum(a, b int) int {
	total := a + b

	return total
}
