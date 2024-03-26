package main

import (
	"fmt"
)

type MyEvent struct {
	Username string `json:"username"`
}

func HandleRequest(event MyEvent) (string, error) {
	if event.Username == "" {
		return "", fmt.Errorf("Username cannot be empty")
	}

	return fmt.Sprintf("Successfully called by %s", event.Username), nil
}

func main() {
	// lambda.Start(HandleRequest)
	fmt.Println("Hello, World!")
}
