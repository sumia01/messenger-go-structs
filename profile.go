package messenger

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Profile struct holds data associated with Facebook profile
type Profile struct {
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	ProfilePicture string `json:"profile_pic,omitempty"`
	Locale         string `json:"locale,omitempty"`
	Timezone       int    `json:"timezone,omitempty"`
	Gender         string `json:"gender,omitempty"`
}

// GetProfile fetches the recipient's profile from facebook platform
// Non empty UserID has to be specified in order to receive the information
func (p *Profile) GetProfile(userID string, accessToken string, url string) error {
	parameters := "fields=first_name,last_name,profile_pic,locale,timezone,gender"
	if url == "" {
		url = fmt.Sprintf("%v/%v/%v?%v&access_token=%v", GraphAPI, GraphAPIVersion, userID, parameters, accessToken)
	} else {
		url = fmt.Sprintf(url+"/%v?%v&access_token=%v", userID, parameters, accessToken)
	}
	resp, err := doRequest("GET", url, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	read, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		er := new(rawError)
		json.Unmarshal(read, er)
		return errors.New("Error occured: " + er.Error.Message)
	}
	profile := Profile{}
	err = json.Unmarshal(read, &profile)
	if err == nil {
		p.FirstName = profile.FirstName
		p.LastName = profile.LastName
		p.ProfilePicture = profile.ProfilePicture
		p.Locale = profile.Locale
		p.Timezone = profile.Timezone
		p.Gender = profile.Gender

	}
	return err
}

type accountLinking struct {
	//Recipient is Page Scoped ID
	Recipient string `json:"recipient"`
}
