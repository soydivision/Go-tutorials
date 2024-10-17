package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("please, provide name")
	}

	message := fmt.Sprintf(randomFormat(), name)
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

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
		"Hello, %v! How are you?",
		"Greetings, %v! Nice to meet you.",
		"Salutations, %v! Hope you're having a great day.",
		"Hey there, %v! What's new?",
		"Welcome aboard, %v! Enjoy your stay.",
	}

	return formats[rand.Intn(len(formats))] + "\n"

}
