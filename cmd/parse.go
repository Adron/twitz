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
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"strings"
)

// parseCmd represents the parse command
var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "This command will extract the Twitter Accounts form a text file.",
	Long: `This command will extract the Twitter Accounts and clean up or disregard other characters 
or text around the twitter accounts to create a simple, clean, Twitter Accounts only list.`,
	Run: func(cmd *cobra.Command, args []string) {
		completedTwittererList := buildTwitterList(false)
		fmt.Println(completedTwittererList)

		willExport := viper.GetString("fileExport")

		if len(willExport) > 1 {
			exportParsedTwitterList(viper.GetString("fileExport"), viper.GetString("fileFormat"), completedTwittererList)
		}
	},
}

func exportParsedTwitterList(exportFilename string, exportFormat string, twittererList []string) {
	if exportFormat == "txt" {
		exportTxt(exportFilename, twittererList, exportFormat)
	} else if exportFormat == "json" {
		exportJson(exportFilename, twittererList, exportFormat)
	} else if exportFormat == "xml" {
		exportXml(exportFilename, twittererList, exportFormat)
	} else if exportFormat == "csv" {
		exportCsv(exportFilename, twittererList, exportFormat)
	} else {
		fmt.Println("Export type unsupported.")
	}
}

func exportXml(exportFilename string, twittererList []string, exportFormat string) {
	fmt.Printf("Starting xml export to %s.", exportFilename)
	xmlContent, err := xml.Marshal(twittererList)
	check(err)
	header := xml.Header
	collectedContent := header + string(xmlContent)
	exportFile(collectedContent, exportFilename+"."+exportFormat)
}

func exportCsv(exportFilename string, twittererList []string, exportFormat string) {
	fmt.Printf("Starting txt export to %s.", exportFilename)
	collectedContent := rebuildForExport(twittererList, ",")
	exportFile(collectedContent, exportFilename+"."+exportFormat)
}

func exportTxt(exportFilename string, twittererList []string, exportFormat string) {
	fmt.Printf("Starting %s export to %s.", exportFormat, exportFilename)
	collectedContent := rebuildForExport(twittererList, "\n")
	exportFile(collectedContent, exportFilename+"."+exportFormat)
}

func exportJson(exportFilename string, twittererList []string, exportFormat string) {
	fmt.Printf("Starting %s export to %s.", exportFormat, exportFilename)
	collectedContent := collectContent(twittererList)
	exportFile(string(collectedContent), exportFilename+"."+exportFormat)
}

func collectContent(twittererList []string) []byte {
	collectedContent, err := json.Marshal(twittererList)
	check(err)
	return collectedContent
}

func rebuildForExport(twittererList []string, concat string) string {
	var collectedContent string
	for _, twitterAccount := range twittererList {
		collectedContent = collectedContent + concat + twitterAccount
	}
	if concat == "," {
		collectedContent = strings.TrimLeft(collectedContent, concat)
	}
	return collectedContent
}

func exportFile(collectedContent string, exportFile string) {
	contentBytes := []byte(collectedContent)
	err := ioutil.WriteFile(exportFile, contentBytes, 0644)
	check(err)
}

func init() {
	rootCmd.AddCommand(parseCmd)
}
