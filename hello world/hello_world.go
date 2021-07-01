package main

import (
	"fmt"

	"log"

	"example.com/error_handling"
	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Suzanne")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

	log.SetPrefix("error_handling: ")

	names := []string{"Hyungsuk", "John", "Stayman"}

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)

	message2, err := error_handling.HelloWithErrorHandling("")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message2)
}
