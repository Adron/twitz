package coreTwitz

import (
	"context"
	"github.com/Adron/twitz/helpers"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

func GetTwitterClient() twitter.Client {
	accessToken, err := helpers.GetBearerToken(viper.GetString("api_key"), viper.GetString("api_secret"))
	helpers.Check(err)
	config := &oauth2.Config{}
	token := &oauth2.Token{AccessToken: accessToken}
	httpClient := config.Client(context.Background(), token)
	client := twitter.NewClient(httpClient)
	return *client
}

