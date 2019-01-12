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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "A quick view into the environment and configuration variables and validation of.",
	Long:  `This command provides a quick view of the configuration variables of the configuration file and any environment variables, and validation of.`,
	Run: func(cmd *cobra.Command, args []string) {
		passed, err := validateRequiredConfig()
		fmt.Printf("Did validation pass? %t\n%s", passed, err)

		if passed {
			fmt.Printf("Twitterers File: %s\n", viper.GetString("file"))
			fmt.Printf("Export File: %s\n", viper.GetString("fileExport"))
			fmt.Printf("Export Format: %s\n", viper.GetString("fileFormat"))
			fmt.Printf("Consumer API Key: %s\n", viper.GetString("consumer_api_key")[0:6])
			fmt.Printf("Consumer API Secret: %s\n", viper.GetString("consumer_api_secret")[0:6])
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
