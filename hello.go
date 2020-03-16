package main

import (
	"fmt"
	"time"
)

func main() {
	counter := 1

	for counter == 1 {
		greeting := hello("Kean Pang")
		fmt.Println(greeting)

		time.Sleep(2 * time.Second)

		total := sum(4, 5)
		fmt.Println(total)

		time.Sleep(2 * time.Second)
	}
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
