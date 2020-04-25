package oauth

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	mux             *http.ServeMux
	server          *httptest.Server
	tokenHandleFunc = func() {
		mux.HandleFunc(tokenEndpoint, func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, `{
			"access_token":"1000.2deaf8d0c268e3c85daa2a013a843b10.703adef2bb337b 8ca36cfc5d7b83cf24",
			"expires_in":3600,
			"api_domain":"https://www.zohoapis.com",
			"token_type":"Bearer"
		}`)
		})
	}
)

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
}

func teardown() {
	server.Close()
}

func TestGetAuthToken_TokenNotExpired(t *testing.T) {
	setup()
	defer teardown()
	tokenHandleFunc()

	authenticator, _ := NewAuthenticator("fakeClientID", "fakeClientSecret", "fakeRefreshToken", server.URL)
	got, _ := authenticator.GetAuthToken()
	want := "1000.2deaf8d0c268e3c85daa2a013a843b10.703adef2bb337b 8ca36cfc5d7b83cf24"
	if got != want {
		t.Errorf("failed getting initial token: expected %v; got %v", want, got)
	}

	// invoke GetAuthToken again to make sure we do not refresh a valid token
	got, _ = authenticator.GetAuthToken()
	if got != want {
		t.Errorf("failed getting token with expiry set: expected %v; got %v", want, got)
	}
}

func TestGetAuthToken_TokenExpired(t *testing.T) {
	setup()
	defer teardown()
	tokenHandleFunc()

	authenticator, _ := NewAuthenticator("fakeClientID", "fakeClientSecret", "fakeRefreshToken", server.URL)
	authenticator.Token.expiry = 1
	got, _ := authenticator.GetAuthToken()
	want := "1000.2deaf8d0c268e3c85daa2a013a843b10.703adef2bb337b 8ca36cfc5d7b83cf24"
	if got != want {
		t.Errorf("failed getting token with expiry set: expected %v; got %v", want, got)
	}
}
