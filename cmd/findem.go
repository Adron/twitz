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
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/spf13/cobra"
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
		for _, account := range completedTwitterList {
			fmt.Println("Looking into %s's TWitter Activity.", account)
			getAccountInfo(account)
		}

	},
}

func getAccountInfo(account string){

	keysTokens := getKeysAndTokens()
	fmt.Println(account)
	fmt.Println(keysTokens)

	consumerKey := keysTokens.ConsumerApiKey
	consumerSecret := keysTokens.ConsumerApiSecret
	accessToken := keysTokens.AccessToken
	accessSecret := keysTokens.AccessTokenSecret

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	// OAuth1 http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	// Verify Credentials
	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}

	user, _, err := client.Accounts.VerifyCredentials(verifyParams)
	check(err)
	fmt.Printf("User's ACCOUNT:\n%+v\n", user)
}

func init() {
	rootCmd.AddCommand(findemCmd)
}
