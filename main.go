package main

import (
	"fmt"
)

func main() {
	fmt.Println ("Hello, World!")
}

// Greeting takes a string and returns a greeting message.
func Greeting(s string) string {
    return "Hello, " + s + "!"
}

// AppendElement takes a slice of integers and an integer value,
// and returns a new slice with the value appended to the end.
func AppendElement(slice []int, i int) []int {
    return append(slice, i)
}