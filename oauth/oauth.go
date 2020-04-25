package oauth

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const (
	tokenEndpoint = "/oauth/v2/token"
)

// Authenticator contains all required authentication information
// to get an access token.
type Authenticator struct {
	ClientID     string
	ClientSecret string
	RefreshToken string
	TokenDomain  string
	Token        Token
}

//Token contains the access token and related information.
type Token struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	expiry      int64
}

// NewAuthenticator returns an Authenticator that contains a
// current access token, which is refreshed automatically based
// on the expiry.
func NewAuthenticator(clientID string, clientSecret string, refreshToken string, tokenDomain string) (*Authenticator, error) {

	authenticator := Authenticator{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RefreshToken: refreshToken,
		TokenDomain:  tokenDomain,
	}

	return &authenticator, nil
}

// getToken fetches an access token through the site24x7 token endpoint
func (auth *Authenticator) GetAuthToken() (string, error) {

	if !isValidToken(auth) {
		if err := refreshToken(auth); err != nil {
			return "", errors.Wrap(err, "cannot refresh token")
		}
	}

	return auth.Token.AccessToken, nil
}

func isValidToken(auth *Authenticator) bool {
	// no token set (initial)
	if auth.Token.expiry == 0 {
		return false
	}
	// token needs to be refreshed
	if auth.Token.expiry-time.Now().Unix() < 500 {
		return false
	}
	// token valid
	return true
}

func refreshToken(auth *Authenticator) error {
	payload := strings.NewReader(fmt.Sprintf("grant_type=refresh_token&client_id=%s&client_secret=%s&refresh_token=%s", auth.ClientID, auth.ClientSecret, auth.RefreshToken))
	req, _ := http.NewRequest("POST", auth.TokenDomain+tokenEndpoint, payload)
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	if strings.Contains(string(body), "error") {
		return errors.New(string(body))
	}

	var token Token
	if err := json.Unmarshal(body, &token); err != nil {
		return errors.Wrapf(err, "cannot unmarshal json")
	}

	auth.Token = token
	epoch := time.Now().Unix()
	auth.Token.expiry = epoch + auth.Token.ExpiresIn

	return nil
}
