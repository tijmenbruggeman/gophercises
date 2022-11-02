package main

import (
	"create-module/greetings"
	"fmt"
	"log"
)

func main() {
    // Get a greeting message and print it.
    message, err := greetings.Hello("Tijmen")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(message)
}