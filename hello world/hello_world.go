package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Suzanne")
	fmt.Println(message)
}
