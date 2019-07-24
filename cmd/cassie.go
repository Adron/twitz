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
	"github.com/Adron/twitz/coreTwitz"
	"github.com/Adron/twitz/storage"
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var cassieCmd = &cobra.Command{
	Use:   "cassie",
	Short: "This command executes parse and inserts the results into Apache Cassandra.",
	Long:  `This command executes parse and inserts the results into Apache Cassandra based on an existing keyspace of "twitz" and respective tables: "accounts". For more details check out the database documentation @ https://github.com/Adron/twitz/database.md for the command.`,
	Run: func(cmd *cobra.Command, args []string) {
		cassieSession := storage.GetCassieSession()
		defer cassieSession.Close()

		twitterList := coreTwitz.BuildTwitterList(true)
		twitterClient := coreTwitz.GetTwitterClient()
		twitterUsers := coreTwitz.GetTwitterDetails(twitterClient, twitterList)

		for _, twitterUser := range twitterUsers {
			fmt.Printf("\nUser: %s", twitterUser.ScreenName)

			storage.InsertTwitterAccount(cassieSession, twitterUser)
		}
	},
}

func init() {
	rootCmd.AddCommand(cassieCmd)
}
