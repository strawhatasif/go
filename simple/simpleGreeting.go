package main

import (
    "fmt"
    "fun.com/greetings"
)

func main() {
    //Retrieve a message and display
    message := greetings.SimpleGreeting("Rachel")
    fmt.Println(message)
}