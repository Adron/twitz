package helpers

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
)

func Check(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func ValidateRequiredConfig() (bool, error) {
	allKeys := viper.AllKeys()
	fmt.Println(allKeys)

	keysPass := true
	errorsList := []string{}

	fileSet := Contains(allKeys, "file")
	consumerApiKeySet := Contains(allKeys, "api_key")
	consumerApiSecretSet := Contains(allKeys, "api_secret")

	if !fileSet {
		errorsList = append(errorsList, "The file to parse for the Twitter list needs to be specified.")
		keysPass = false
	}

	if !consumerApiKeySet || !consumerApiSecretSet {
		errorsList = append(errorsList, "Required API key and secret configuration variables are not set.")
		keysPass = false
	}

	if consumerApiKeySet && consumerApiSecretSet {
		apiKey := viper.GetString("api_key")
		apiSecret := viper.GetString("api_secret")

		if len(apiKey) == 0 || len(apiSecret) == 0 {
			errorsList = append(errorsList, "The API Key and Secret have a zero length and appear to be empty strings. Please set the TWITZ_API_KEY and TWITZ_API_SECRET environment variables.")
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
