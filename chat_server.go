package main

import (
	"fmt"

	. "github.com/pspaces/gospace"
)

func main() {
	chat := NewSpace("tcp://localhost:31415/master")
	var who string
	var message string

	for {
		t, _ := chat.QueryAll(&who, &message)
		who = (t.GetFieldAt(0)).(string)
		message = (t.GetFieldAt(1)).(string)
		fmt.Printf("%s: %s \n", who, message)
		fmt.Printf(t)
	}
}
