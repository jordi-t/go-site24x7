package site24x7

import (
	"encoding/json"
	"io/ioutil"
)

// UserService provides an interface Site24x7 user management.
type UserService struct {
	client *client
}

// Users contains a list of Site24x7 users.
type UserList struct {
	Users []User `json:"data"`
}

// User represents a Site24x7 user.
type User struct {
	ImagePresent    bool `json:"image_present"`
	TwitterSettings struct {
	} `json:"twitter_settings"`
	SelectionType    int    `json:"selection_type"`
	EmailAddress     string `json:"email_address"`
	IsAccountContact bool   `json:"is_account_contact"`
	IsContact        bool   `json:"is_contact"`
	AlertSettings    struct {
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
	} `json:"alert_settings"`
	UserGroups []string `json:"user_groups"`
	IsInvited  bool     `json:"is_invited"`
	ImSettings struct {
	} `json:"im_settings"`
	NotifyMedium   []int  `json:"notify_medium"`
	IsEditAllowed  bool   `json:"is_edit_allowed"`
	DisplayName    string `json:"display_name"`
	UserID         string `json:"user_id"`
	MobileSettings struct {
		CountryCode    string `json:"country_code"`
		SmsProviderID  int    `json:"sms_provider_id"`
		MobileNumber   string `json:"mobile_number"`
		CallProviderID int    `json:"call_provider_id"`
	} `json:"mobile_settings"`
	UserRole int    `json:"user_role"`
	JobTitle int    `json:"job_title"`
	Zuid     string `json:"zuid"`
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
