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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// parseCmd represents the parse command
var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "This command will extract the Twitter Accounts form a text file.",
	Long: `This command will extract the Twitter Accounts and clean up or disregard other characters 
or text around the twitter accounts to create a simple, clean, Twitter Accounts only list.`,
	Run: func(cmd *cobra.Command, args []string) {
		completedTwittererList := coreTwitz.BuildTwitterList(false)
		fmt.Println(completedTwittererList)

		willExport := viper.GetString("fileExport")

		if len(willExport) > 1 {
			exportParsedTwitterList(viper.GetString("fileExport"), viper.GetString("fileFormat"), completedTwittererList)
		}
	},
}

func exportParsedTwitterList(exportFilename string, exportFormat string, twittererList []string) {
	if exportFormat == "txt" {
		coreTwitz.ExportTxt(exportFilename, twittererList, exportFormat)
	} else if exportFormat == "json" {
		coreTwitz.ExportJson(exportFilename, twittererList, exportFormat)
	} else if exportFormat == "xml" {
		coreTwitz.ExportXml(exportFilename, twittererList, exportFormat)
	} else if exportFormat == "csv" {
		coreTwitz.ExportCsv(exportFilename, twittererList, exportFormat)
	} else {
		fmt.Println("Export type unsupported.")
	}
}

func init() {
	rootCmd.AddCommand(parseCmd)
}
