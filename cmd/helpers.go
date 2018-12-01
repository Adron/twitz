package cmd

import (
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
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

type keysAndTokens struct {
	ConsumerApiKey, ConsumerApiSecret, AccessToken, AccessTokenSecret string
}

func getKeysAndTokens() keysAndTokens {
	keysTokens := keysAndTokens{
		viper.GetString("consumer_api_key"),
		viper.GetString("consumer_api_secret"),
		viper.GetString("access_token"),
		viper.GetString("access_token_secret")}
	return keysTokens
}