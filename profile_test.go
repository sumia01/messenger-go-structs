package messenger

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

func TestGetProfile(t *testing.T) {
	//Avoid HTTPS in tests
	GraphAPI = "http://example.com"

	mockData := &Profile{
		FirstName:      "John",
		LastName:       "Smith",
		ProfilePicture: "https://example.com/",
		Gender:         "male",
		Timezone:       -5,
		Locale:         "en_US",
	}

	body, err := json.Marshal(mockData)
	if err != nil {
		t.Error(err)
	}

	setClient(200, body)
	prof := Profile{}
	err = prof.GetProfile("123", "321", "")
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(&prof, mockData) {
		t.Error("Profiles do not match")
	}

	errorData := &rawError{Error: Error{
		Message: "w/e",
	}}
	body, err = json.Marshal(errorData)
	if err != nil {
		t.Error(err)
	}
	setClient(400, body)
	prof2 := Profile{}
	err = prof2.GetProfile("123", "321", "")
	if err.Error() != "Error occured: "+errorData.Error.Message {
		t.Error("Invalid error parsing")
	}
}

func setClient(code int, body []byte) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(code)
		w.Write(body)
	}))

	http.DefaultClient.Transport = &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}
	return server
}
