package coreTwitz

import (
	"fmt"
	"github.com/Adron/twitz/helpers"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/satori/go.uuid"
	"github.com/spf13/viper"
	"io/ioutil"
)

const fileExport = "fileExport"
const fileFormat = "fileFormat"

type TwitterDerived struct {
	TwitterAccounts []twitter.User
}

type TwitterParsed struct {
	TwitterNames []string
}

type PrinterAndExport interface {
	PrintAndExport()
}

func (tDerived TwitterDerived) PrintAndExport() {
	exportFilename := viper.GetString(fileExport)
	exportFormat := viper.GetString(fileFormat)

	if len(exportFilename) > 1 {
		if exportFormat == "txt" {
			ExportTwitterUsersTxt(tDerived.TwitterAccounts)
		} else if exportFormat == "json" {
			ExportTwitterUsersJson(tDerived.TwitterAccounts)
		} else if exportFormat == "xml" {
			ExportTwitterUsersXml(tDerived.TwitterAccounts)
		} else if exportFormat == "csv" {
			ExportTwitterUsersCsv(tDerived.TwitterAccounts)
		} else {
			fmt.Println("Export type unsupported.")
		}
	}
}

func (tParsed TwitterParsed) PrintAndExport() {
	exportFilename := viper.GetString(fileExport)
	exportFormat := viper.GetString(fileFormat)

	if exportFormat == "txt" {
		ExportTxt(exportFilename, tParsed.TwitterNames, exportFormat)
	} else if exportFormat == "json" {
		ExportJson(exportFilename, tParsed.TwitterNames, exportFormat)
	} else if exportFormat == "xml" {
		ExportXml(exportFilename, tParsed.TwitterNames, exportFormat)
	} else if exportFormat == "csv" {
		ExportCsv(exportFilename, tParsed.TwitterNames, exportFormat)
	} else {
		fmt.Println("Export type unsupported.")
	}
}

func (tDerived TwitterDerived) FileExporter(twitterList string) {
	exportFile(twitterList)
}

func (tParsed TwitterParsed) FileExporter(twitterList string) {
	exportFile(twitterList)
}

func exportFile(collectedContent string) {
	if viper.GetString(fileExport) != "" {
		contentBytes := []byte(collectedContent)
		writeToFile := viper.GetString(fileExport) +
			uuid.NewV4().String() + "." +
			viper.GetString(fileFormat)
		err := ioutil.WriteFile(writeToFile, contentBytes, 0644)
		helpers.Check(err)
	}
}

func GetTwitterDetails(client twitter.Client, list []string) []twitter.User {
	userLookupParams := &twitter.UserLookupParams{ScreenName: list}
	twitterUsers, _, _ := client.Users.Lookup(userLookupParams)
	return twitterUsers
}

