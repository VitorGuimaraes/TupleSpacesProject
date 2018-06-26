package main

import (
	"fmt"

	. "github.com/pspaces/gospace"
)

func main() {
	chat := NewRemoteSpace("tcp://localhost:31415/master")
	_, err := chat.Put("Alice", "Hi!")

	if err != nil {
		fmt.Printf("Something went wrong when trying to send a message: %v\n", err)
	}
}
