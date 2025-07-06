package main

import (
	err "errors" // err is an alias of errors package;
	f "fmt"
)

func main() {
	f.Println("Hello World!")
}

func sayHello(name string) (string, error) {
	if name == "" {
		return "", err.New("no name provided")
	}
	return "Hello " + name, nil
}