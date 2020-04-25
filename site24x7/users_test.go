package site24x7

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{
				 "code": 0,
				 "message": "success",
				 "data": [{
				   "email_address": "joe@company.com"
				 },{
				   "email_address": "jane@company.com"
				 }]
				}`)
	})

	want := UserList{
		Users: []User{
			{
				EmailAddress: "joe@company.com",
			},
			{
				EmailAddress: "jane@company.com",
			},
		},
	}

	m := &MockAuthenticator{}
	c := NewClient(m, server.URL)
	got, _ := c.Users.List()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("failed: expected %v; got %v", want, got)
	}
}
