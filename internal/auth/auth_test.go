package auth

import (
	"net/http"
	"testing"
)

func TestValidAPIKEY(t *testing.T) {
	header := http.Header{}
	header.Add("Authorization", "ApiKey 1234")

	got, err := GetAPIKey(header)

	if got != "1234" || err != nil {
		t.Errorf("Error getting API Key\t Received%v\n", got)
	}
}

func TestInvalidAPIKey(t *testing.T) {
	header := http.Header{}
	header.Add("Authorization", " Invalid ApiKey")

	_, err := GetAPIKey(header)

	if err.Error() != "malformed authorization header" {
		t.Errorf("Received unexpected err:\t%v", err.Error())
	}
}

func TestNoAuthHeader(t *testing.T) {
	header := http.Header{}
	header.Add("header1", "Cool Header")

	_, err := GetAPIKey(header)

	if err.Error() == "no authorization header included" {
		t.Errorf("Received unexpected err:\t%v", err.Error())
	}
}
