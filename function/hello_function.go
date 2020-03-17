package function

import "fmt"

//Hello function to return a greeting sentence to the user
func Hello(user string) string {
	if len(user) == 0 {
		return "Hello Dude!"
	}

	return fmt.Sprintf("Hello %v", user)
}

//Sum function to return the sum of two number
func Sum(a, b int) int {
	total := a + b

	return total
}
