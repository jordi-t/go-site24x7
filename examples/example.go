package main

import (
	"fmt"
	"github.com/jseris/go-site24x7/oauth"
	"github.com/jseris/go-site24x7/site24x7"
	"os"
)

func main() {
	clientID := os.Getenv("CLIENTID")
	clientSecret := os.Getenv("CLIENTSECRET")
	refreshToken := os.Getenv("REFRESHTOKEN")
	tokenDomain := os.Getenv("TOKENDOMAIN")
	apiDomain := os.Getenv("APIDOMAIN")

	// Get a new Authenticator
	auth, err := oauth.NewAuthenticator(clientID, clientSecret, refreshToken, tokenDomain)
	if err != nil {
		fmt.Println(err)
	}

	// Get a new site24x7 client
	client := site24x7.NewClient(auth, apiDomain)
	users, err := client.Users.List()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", users)
}
