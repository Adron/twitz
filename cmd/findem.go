// Copyright © 2018 Adron Hall <adron@thrashingcode.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"context"
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

// findemCmd represents the findem command
var findemCmd = &cobra.Command{
	Use:   "findem",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		// TODO: Work through each account and retrieve key information displaying it to screen and exporting to file.

		fmt.Println("Starting Twitter Information Retrieval.")
		completedTwitterList := buildTwitterList()

		fmt.Println(completedTwitterList)

		accessToken, err := getToken(viper.GetString("consumer_api_key"), viper.GetString("consumer_api_secret"))
		check(err)

		fmt.Printf("Access Token Retreived: %s", accessToken)

		config := &oauth2.Config{}
		token := &oauth2.Token{AccessToken: accessToken}
		httpClient := config.Client(context.Background(), token)
		client := twitter.NewClient(httpClient)

		userParams := &twitter.UserLookupParams{ScreenName: []string{"@Adron"}}
		users, _, err := client.Users.Lookup(userParams)

		fmt.Printf("\n\nUser Info: \n%+v\n", users)

		howManyUsersFound := len(users)
		fmt.Println(howManyUsersFound)
	},
}

func init() {
	rootCmd.AddCommand(findemCmd)
}
