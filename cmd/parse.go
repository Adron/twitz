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
	"github.com/Adron/twitz/coreTwitz"
	"github.com/spf13/cobra"
)

// parseCmd represents the parse command
var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "This command will extract the Twitter Accounts form a text file.",
	Long: `This command will extract the Twitter Accounts and clean up or disregard other characters 
or text around the twitter accounts to create a simple, clean, Twitter Accounts only list.`,
	Run: func(cmd *cobra.Command, args []string) {
		completedTwittererList := coreTwitz.BuildTwitterList(false)
		var p = coreTwitz.TwitterParsed{TwitterNames: completedTwittererList}
		coreTwitz.ProcessTwitterAccounts(p)
	},
}

func init() {
	rootCmd.AddCommand(parseCmd)
}
