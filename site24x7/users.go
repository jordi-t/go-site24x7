package site24x7

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"strings"
)

// UserService provides an interface Site24x7 user management.
type UserService struct {
	client *Client
}

// Users contains a list of Site24x7 users.
type UserList struct {
	Users []User `json:"data"`
}

// User represents a Site24x7 user.
type User struct {
	ImagePresent     bool           `json:"image_present,omitempty"`
	SelectionType    int            `json:"selection_type,omitempty"`
	EmailAddress     string         `json:"email_address"`
	IsAccountContact bool           `json:"is_account_contact,omitempty"`
	IsContact        bool           `json:"is_contact,omitempty"`
	AlertSettings    AlertSettings  `json:"alert_settings"`
	UserGroups       []string       `json:"user_groups,omitempty"`
	IsInvited        bool           `json:"is_invited,omitempty"`
	NotifyMedium     []int          `json:"notify_medium,omitempty"`
	IsEditAllowed    bool           `json:"is_edit_allowed,omitempty"`
	DisplayName      string         `json:"display_name"`
	UserID           string         `json:"user_id,omitempty"`
	MobileSettings   MobileSettings `json:"mobile_settings"`
	UserRole         int            `json:"user_role,omitempty"`
	JobTitle         int            `json:"job_title,omitempty"`
	Zuid             string         `json:"zuid,omitempty"`
}

type AlertSettings struct {
	Trouble         []int `json:"trouble"`
	Up              []int `json:"up"`
	DontAlertOnDays []int `json:"dont_alert_on_days"`
	EmailFormat     int   `json:"email_format"`
	AlertingPeriod  struct {
		EndTime   string `json:"end_time"`
		StartTime string `json:"start_time"`
	} `json:"alerting_period"`
	Down    []int `json:"down"`
	Applogs []int `json:"applogs"`
	Anomaly []int `json:"anomaly"`
}

type MobileSettings struct {
	CountryCode    string `json:"country_code"`
	SmsProviderID  int    `json:"sms_provider_id,omitempty"`
	MobileNumber   string `json:"mobile_number,omitempty"`
	CallProviderID int    `json:"call_provider_id"`
}

// List returns all users in the Site24x7 account.
func (us *UserService) List() (UserList, error) {
	res, err := us.client.get("/users")
	if err != nil {
		return UserList{}, err
	}

	body, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	users := UserList{}
	if err := json.Unmarshal(body, &users); err != nil {
		return UserList{}, err
	}

	return users, nil
}

// Create a user
func (us *UserService) Create(user User) (User, error) {
	b, err := json.Marshal(user)
	if err != nil {
		return User{}, errors.Wrap(err, "cannot marshal the User struct")
	}

	res, err := us.client.post("/users", b)
	if err != nil {
		return User{}, errors.Wrap(err, "post request failed")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return User{}, errors.Wrap(err, "cannot read response stream")
	}
	defer res.Body.Close()

	if strings.Contains(string(body), "error") {
		return User{}, errors.New(string(body))
	}

	userList := UserList{}
	if err := json.Unmarshal(body, &userList); err != nil {
		return User{}, err
	}

	fmt.Println(userList.Users[0])

	return userList.Users[0], nil
}
