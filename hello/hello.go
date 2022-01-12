package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Dipankar Das")
	fmt.Println(message)
}
