package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// This function takes a name parameter whose type is string
// This function also returns a string
// In Go, a function that starts with a capital letter can be called by a function not in the same package (exported name)
func Hello(name string) (string, error) {
	if name == "" {
		return name, errors.New("empty name")
	}

	message := fmt.Sprintf(randomFormat(), name)
	//message := fmt.Sprintf(randomFormat())
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		messages[name] = message
	}
	return messages, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages.
// The returned message is selected at random.
func randomFormat() string {
	// A slice of message formats
	formats := []string{
		"Hi, %v. Welcome!",
		"Halo, %v, Willkommen!",
		"안녕하세요, %v, 반갑습니다!",
		"你好，%v, 欢迎欢迎",
	}

	return formats[rand.Intn(len(formats))]
}
