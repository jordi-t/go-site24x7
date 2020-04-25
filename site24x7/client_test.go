package site24x7

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var (
	mux    *http.ServeMux
	server *httptest.Server
)

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

}

func teardown() {
	server.Close()
}

type MockAuthenticator struct {
}

func (mock *MockAuthenticator) GetAuthToken() (string, error) {
	return "fakeAuthToken", nil
}

func TestGet(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"test":"ok"}`)
	})

	m := &MockAuthenticator{}
	c := NewClient(m, server.URL)
	want := `{"test":"ok"}`

	res, _ := c.get("/test")
	body, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	if string(body) != want {
		t.Errorf("failed: expected %v; got %v", want, string(body))
	}

	// use custom http client
	c = NewClient(m, server.URL, http.Client{Timeout: time.Second * 10})
	res, _ = c.get("/test")
	body, _ = ioutil.ReadAll(res.Body)
	if string(body) != want {
		t.Errorf("failed: expected %v; got %v", want, string(body))
	}
}
