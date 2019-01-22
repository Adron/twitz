package cmd

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func buildTwitterList() []string {
	theFile := viper.GetString("file")
	theTwitterers, err := ioutil.ReadFile(theFile)
	check(err)
	stringTwitterers := string(theTwitterers[:])
	splitFields := strings.Fields(stringTwitterers)
	var completedTwittererList []string
	for _, aField := range splitFields {
		if strings.HasPrefix(aField, "@") && aField != "@" {
			reg, _ := regexp.Compile("[^a-zA-Z0-9_@]")
			processedString := reg.ReplaceAllString(aField, "")
			completedTwittererList = append(completedTwittererList, processedString)
		}
	}
	return completedTwittererList
}

func getBearerToken(consumerKey, consumerSecret string) (string, error) {
	req, err := http.NewRequest("POST", "https://api.twitter.com/oauth2/token",
		strings.NewReader("grant_type=client_credentials"))

	if err != nil {
		return "", fmt.Errorf("cannot create /token request: %+v", err)
	}

	b64Token := base64.StdEncoding.EncodeToString(
		[]byte(fmt.Sprintf("%s:%s", consumerKey, consumerSecret)))
	req.Header.Add("Authorization", "Basic "+b64Token)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")

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

func validateRequiredConfig() (bool, error) {
	allKeys := viper.AllKeys()
	fmt.Println(allKeys)

	keysPass := true
	errorsList := []string{}

	fileFormatSet := Contains(allKeys, "fileformat")
	fileSet := Contains(allKeys, "file")
	fileExportSet := Contains(allKeys, "fileexport")
	consumerApiKeySet := Contains(allKeys, "api_key")
	consumerApiSecretSet := Contains(allKeys, "api_secret")

	if !fileFormatSet || !fileSet || !fileExportSet {
		errorsList = append(errorsList, "Required configuration for export are not set. File format, file, and export values need set in the configuration file.")
		keysPass = false
	}

	if !consumerApiKeySet || !consumerApiSecretSet {
		errorsList = append(errorsList, "Required API key and secret configuration variables are not set.")
		keysPass = false
	}

	if consumerApiKeySet && consumerApiSecretSet {
		apiKey := viper.GetString("API_KEY")
		apiSecret := viper.GetString("API_SECRET")

		if len(apiKey) == 0 || len(apiSecret) == 0 {
			errorsList = append(errorsList, "The API Key and Secret have a zero length and appear to be empty strings.")
			keysPass = false
		} else if len(apiKey) < 6 || len(apiSecret) < 6 {
			errorsList = append(errorsList, "The key or secret key also appear to be malformed. Please verify and enter a correct API Key and Secret.")
			keysPass = false
		}
	}

	var collectedErrors = "Errors:"
	for _, errorItem := range errorsList {
		collectedErrors = collectedErrors + "\n" + errorItem
	}
	configurationError := errors.New(collectedErrors)

	return keysPass, configurationError
}

func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
