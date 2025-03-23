package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello retruns a greeting for the named person.
func Hello(name string) (string, error) {
	// if no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}

	// Return a greeting that embeds the name in a message.
	message := fmt.Sprint(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages.
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)

		if err != nil {
			return nil, err
		}

		// in the map, associate the retrieved message with
		// the name.
		messages[name] = message

	}
	return messages, nil
}

// randomFormat return one of a set of greeting messages returned
// message is selected at random
func randomFormat() string {
	// A slice of a message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Greate to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Return a randomly a selected message format by specifying
	// a random index for the slice formats.
	return formats[rand.Intn(len(formats))]

}
