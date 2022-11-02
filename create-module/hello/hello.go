package main

import (
	"create-module/greetings"
	"fmt"
	"log"
)

func main() {
    // Get a greeting message and print it.
    names := []string{"Tijmen", "Debbie", "Mila", "Bo", "Hidde"}
    message, err := greetings.Hellos(names)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(message)
}