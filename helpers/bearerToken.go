package helpers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func GetBearerToken(consumerKey, consumerSecret string) (string, error) {
	req := buildBearerTokenRequest()
	setHeadersEncodeBase64(consumerKey, consumerSecret, req)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("/token request failed: %+v", err)
	}
	defer resp.Body.Close()

	var v struct {
		AccessToken string `json:"access_token"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
		return "", fmt.Errorf("error parsing json in /token response: %+v", err)
	}
	if v.AccessToken == "" {
		return "", fmt.Errorf("/token response does not have access_token")
	}
	return v.AccessToken, nil
}

func setHeadersEncodeBase64(consumerKey string, consumerSecret string, req *http.Request) {
	b64Token := base64.StdEncoding.EncodeToString(
		[]byte(fmt.Sprintf("%s:%s", consumerKey, consumerSecret)))
	req.Header.Add("Authorization", "Basic "+b64Token)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
}

func buildBearerTokenRequest() *http.Request {
	req, err := http.NewRequest("POST", "https://api.twitter.com/oauth2/token",
		strings.NewReader("grant_type=client_credentials"))
	Check(err)
	return req
}

