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
	"io/ioutil"
	"regexp"
	"strings"
)

// parseCmd represents the parse command
var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "This command will extract the Twitter Accounts form a text file.",
	Long: `This command will extract the Twitter Accounts and clean up or disregard other characters 
or text around the twitter accounts to create a simple, clean, Twitter Accounts only list.`,
	Run: func(cmd *cobra.Command, args []string) {
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

		fmt.Println(completedTwittererList)

		exporterThingy(viper.GetString("fileExport"), viper.GetString("fileFormat"), completedTwittererList)
	},
}

func exporterThingy(exportFilename string, exportFormat string, twittererList []string) {
	if exportFormat == "txt" {
		fmt.Printf("Starting txt export to %s.", exportFilename)

		var collectedContent string
		for _, twitterAccount := range twittererList {
			collectedContent = collectedContent + "\n" + twitterAccount
		}

		contentBytes := []byte(collectedContent)
		err := ioutil.WriteFile(exportFilename+"."+exportFormat, contentBytes, 0644)
		check(err)
	} else if exportFormat == "json" {

	} else if exportFormat == "xml" {

	} else if exportFormat == "csv" {

	} else {
		fmt.Println("Export type unsupported.")
	}
}

func init() {
	rootCmd.AddCommand(parseCmd)
	//rootCmd.PersistentFlags().StringVar(&export, "export", "file-being-exported", "Set this by passing in the export file.")
	//viper.BindPFlag("export", rootCmd.PersistentFlags().Lookup("export"))
}
