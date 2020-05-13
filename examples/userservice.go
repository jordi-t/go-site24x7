package main

import (
	"fmt"
	"github.com/jseris/go-site24x7/oauth"
	"github.com/jseris/go-site24x7/site24x7"
	"os"
)

func createUser(client *site24x7.Client) {

	alertSettings := site24x7.AlertSettings{
		Trouble:         []int{1, 2, 3},
		Up:              []int{1, 2, 3},
		DontAlertOnDays: nil,
		EmailFormat:     0,
		AlertingPeriod: struct {
			EndTime   string `json:"end_time"`
			StartTime string `json:"start_time"`
		}{"20:15", "4:30"},
		Down:    []int{1, 2, 3},
		Applogs: []int{1, 2, 3},
		Anomaly: []int{1},
	}

	mobileSettings := site24x7.MobileSettings{
		CountryCode:  "31",
		MobileNumber: "0612345678",
	}

	user := site24x7.User{
		EmailAddress:   "test@testuser.com",
		AlertSettings:  alertSettings,
		DisplayName:    "Test User",
		UserRole:       1,
		NotifyMedium:   []int{1, 2, 3},
		MobileSettings: mobileSettings,
	}

	b, err := client.Users.Create(user)
	if err != nil {
		fmt.Printf("error: %+v", err)
		return
	}

	fmt.Printf("CreateUser response: %s", string(b))
}

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

	// Create a user
	createUser(client)

	// List users
	users, err := client.Users.List()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", users)
}
