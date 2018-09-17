package main

import (
	"log"
	"os"

	"github.com/Dids/hackchat/api"
)

func main() {
	// TODO: Check for existing token and fallback to password

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
