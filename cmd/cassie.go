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
	"github.com/Adron/twitz/coreTwitz"
	"github.com/Adron/twitz/helpers"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/gocql/gocql"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"log"
)

// configCmd represents the config command
var cassieCmd = &cobra.Command{
	Use:   "cassie",
	Short: "This command executes parse and inserts the results into Apache Cassandra.",
	Long:  `This command executes parse and inserts the results into Apache Cassandra based on an existing keyspace of "twitz" and respective tables: "accounts". For more details check out the database documentation @ https://github.com/Adron/twitz/database.md for the command.`,
	Run: func(cmd *cobra.Command, args []string) {
		cluster := gocql.NewCluster(viper.GetString("cassie"))
		cluster.Keyspace = viper.GetString("keyspace")
		cluster.Consistency = gocql.Quorum
		session, _ := cluster.CreateSession()
		defer session.Close()

		fmt.Println("Starting Twitter Information Retrieval.")
		completedTwittererList := coreTwitz.BuildTwitterList(true)
		fmt.Printf("Getting Twitter details for: \n%s", completedTwittererList)

		accessToken, err := helpers.GetBearerToken(viper.GetString("api_key"), viper.GetString("api_secret"))
		helpers.Check(err)

		config := &oauth2.Config{}
		token := &oauth2.Token{AccessToken: accessToken}
		httpClient := config.Client(context.Background(), token)
		client := twitter.NewClient(httpClient)

		userLookupParams := &twitter.UserLookupParams{ScreenName: completedTwittererList}

		twitterUsers, _, _ := client.Users.Lookup(userLookupParams)

		for _, twitterUser := range twitterUsers {

			if err := session.Query(`INSERT INTO twitz.twitterers (id, handle, name) VALUES (?, ?, ?)`,
				gocql.TimeUUID(), twitterUser.Name, twitterUser.ScreenName).Exec(); err != nil {
				log.Fatal(err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(cassieCmd)
}
