package greetings

import "fmt"

// This function takes a name parameter whose type is string
// This function also returns a string
// In Go, a function that starts with a capital letter can be called by a function not in the same package (exported name)
func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
