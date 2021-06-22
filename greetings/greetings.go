package greetings

import "fmt"

func SimpleGreeting(name string) string {
    //this returns a greeting with the name embedded
    message := fmt.Sprintf("Why hello, %v. a simple hello for you!", name)
    return message
}