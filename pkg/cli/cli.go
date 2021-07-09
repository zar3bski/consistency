package cli

import "fmt"

func greetings() {
	fmt.Println("Hello")
}

func Name() string {
	name := "David"
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
