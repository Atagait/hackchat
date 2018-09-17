package api

import (
	"fmt"
	"log"

	"github.com/imroc/req"
)

// Documentation/references:
// https://hackmud.zendesk.com/hc/en-us/articles/115003767188-1-4-4-patch-notes
// https://hackmud-unofficial.github.io/hackmud-chat/

// BaseURL is a constant path to the API root
const BaseURL = "https://www.hackmud.com/mobile/"

// Endpoint is a simple string type for API endpoint paths
type Endpoint string

const (
	// GetTokenEndpoint returns the chat token
	GetTokenEndpoint Endpoint = "get_token.json"

	// AccountDataEndpoint returns the user's account data
	AccountDataEndpoint Endpoint = "account_data.json"

	// ChatsEndpoint returns TODO
	ChatsEndpoint Endpoint = "chats.json"

	// ChatHistoryEndpoint returns TODO
	ChatHistoryEndpoint Endpoint = "chat_history.json"

	// CreateChatEndpoint returns TODO
	CreateChatEndpoint Endpoint = "create_chat.json"
)

// UserName TODO
type UserName string

// ChannelName TODO
type ChannelName string

// Channel TODO
type Channel map[ChannelName][]UserName

// Users TODO
type Users map[UserName]Channel

// Chat TODO
type Chat struct {
	ID      int64  `json:"id"`
	Time    int64  `json:"t"`
	Channel string `json:"channel"`
	From    string `json:"from_user"` // TODO: This might be wrong
	To      string `json:"to_user"`   // TODO: This might be wrong
	Message string `json:"msg"`
	IsJoin  bool   `json:"is_join"`
	IsLeave bool   `json:"is_leave"`
}

// Chats TODO
type Chats map[UserName][]Chat

// Response is the response object for the API endpoints
type Response struct {
	OK               bool     `json:"ok"`
	ChatToken        string   `json:"chat_token"`
	Message          string   `json:"msg"`
	Users            Users    `json:"users"`
	Chats            Chats    `json:"chats"`
	InvalidUsernames []string `json:"invalidUsernames"`
}

// ChatToken stores the session token ("chat_token")
var ChatToken string

// Login logs you in to the API and stores the "chat_token" variable
func Login(password string) {
	// Notify the user
	fmt.Println("Logging in to the API..")

	// Attempt to login by calling the "get_token" endpoint
	result, err := CallEndpoint(GetTokenEndpoint, req.Param{
		"pass": password,
	})

	// Validate the response
	if err != nil {
		log.Fatal("Failed to log in: ", err)
	} else if !result.OK || len(result.ChatToken) <= 0 {
		if len(result.Message) > 0 {
			log.Fatal("Failed to log in: ", result.Message)
		}
		log.Fatal("Failed to log in: ", "an unknown error occurred")
	}

	// TODO: We should figure out a way to safely store the token instead,
	//       then simply refresh/reload the token when we start again,
	//       so we don't need to (re)login every time..

	// Store the "chat_token"
	ChatToken = result.ChatToken

	// Notify the user
	fmt.Println("Successfully logged in!")
}

// LoadAccountData loads the latest account data associated with the token
func LoadAccountData(token string) {
	// Notify the user
	fmt.Println("Loading account data from the API..")

	// Attempt to login by calling the "account_data" endpoint
	result, err := CallEndpoint(AccountDataEndpoint, req.Param{
		"chat_token": token,
	})

	// Validate the response
	if err != nil {
		log.Fatal("Failed to load account data: ", err)
	} else if !result.OK || len(result.Users) <= 0 {
		if len(result.Message) > 0 {
			log.Fatal("Failed to load account data: ", result.Message)
		}
		log.Fatal("Failed to load account data: ", "an unknown error occurred")
	}

	// TODO: Store the account data locally?

	// Store the "chat_token"
	//ChatToken = result.ChatToken

	// Notify the user
	fmt.Println("Successfully loaded account data!")
}

// FIXME: We need some sample chat data, because there's nothing there at the moment
//        Sample response: {"ok":true,"chats":{"dids":[]}}

// LoadChats loads the chat data associated with the token
func LoadChats(token string) {
	// Notify the user
	fmt.Println("Loading chat data from the API..")

	// Attempt to login by calling the "chats" endpoint
	result, err := CallEndpoint(ChatsEndpoint, req.Param{
		"chat_token": token,
		"usernames":  []string{"some_username"}, // FIXME: Do NOT hardcode this, but instead pass in our users (somehow get the usernames from the map..)
	})

	// Validate the response
	if err != nil {
		log.Fatal("Failed to load chat data: ", err)
	} else if !result.OK || len(result.Users) <= 0 {
		if len(result.Message) > 0 {
			log.Fatal("Failed to load chat data: ", result.Message)
		}
		log.Fatal("Failed to load chat data: ", "an unknown error occurred")
	}

	// TODO: Store the chat data? Or just return it?

	// Store the "chat_token"
	//ChatToken = result.ChatToken

	// Notify the user
	fmt.Println("Successfully loaded chat data!")
}

// CallEndpoint fires a GET request for the supplied API endpoint
func CallEndpoint(endpoint Endpoint, params req.Param) (Response, error) {
	// Construct the full url
	fullURL := BaseURL + endpoint
	//fmt.Println("Calling endpoint:", fullURL)

	// TODO: Don't leave this enabled forever!
	// Enable request debugging
	req.Debug = true

	// Fire the request
	response, requestError := req.Post(string(fullURL), req.BodyJSON(&params))
	if requestError != nil {
		return Response{}, requestError
	}

	// Print the response
	//fmt.Println(response)

	// Parse the response as JSON and map it to our custom struct
	var result Response
	response.ToJSON(&result)

	// TODO: Remove this
	fmt.Println("Parsed response object:", result)

	// Return the response struct
	return result, nil
}
