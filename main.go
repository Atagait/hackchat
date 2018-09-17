package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Dids/hackchat/api"
)

// Version is set dynamically when building
var Version = "0.0.0"

// TODO: Separate the API into it's own directory/library,
//       so it can be easily reused in other projects as well

func main() {
	fmt.Println("Version:", Version)
	fmt.Println()

	// Get the password argument
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatal("Invalid or missing password argument")
	}
	password := args[0]

	// Login to the API
	api.Login(password)

	// Load account data
	api.LoadAccountData(api.ChatToken)

	// TODO: Disabled until we can fix the method itself
	// Load chat data
	//api.LoadChats(api.ChatToken)
}
