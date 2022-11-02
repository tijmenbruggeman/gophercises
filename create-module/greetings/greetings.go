package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func init() {
    rand.Seed(time.Now().UnixNano())
}


// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
    // Return a greeting that embeds the name in a message.
    if name == "" {
        return "", errors.New("empty name")
    }
    format := randGreeting();
    message := fmt.Sprintf(format, name)
    return message, nil
}

func randGreeting() string {
    formats := []string {
        "Morge %v",
        "Wat fijn dat je er bent %v",
        "Moin %v",
    }

    return formats[rand.Intn(len(formats))]
}