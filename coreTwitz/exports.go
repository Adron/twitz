package coreTwitz

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/Adron/twitz/helpers"
	"github.com/dghubble/go-twitter/twitter"
	"strconv"
	"strings"
)

func ExportTwitterUsersTxt(twittererList []twitter.User) {
	fmt.Println(twittererList)
	var twitterCollection string
	printTwitterAccount(twittererList, twitterCollection)
}

func ExportTwitterUsersJson(twittererList []twitter.User) {
	collectedContent, err := json.Marshal(twittererList)
	helpers.Check(err)
	exportFile(string(collectedContent))
}

func ExportTwitterUsersXml(twittererList []twitter.User) {
	collectContent, err := xml.Marshal(twittererList)
	helpers.Check(err)
	header := xml.Header
	collectedContent := header + string(collectContent)
	exportFile(collectedContent)
}

func ExportTwitterUsersCsv(twittererList []twitter.User) {
	fmt.Println(twittererList)
}

func ExportCsv(exportFilename string, twittererList []string, exportFormat string) {
	ExportCsvTxtWorkload(exportFilename, twittererList, exportFormat, ",")
}

func ExportTxt(exportFilename string, twittererList []string, exportFormat string) {
	ExportCsvTxtWorkload(exportFilename, twittererList, exportFormat, " ")
}

func ExportCsvTxtWorkload(exportFilename string, twittererList []string, exportFormat string, concat string) {
	fmt.Printf("Starting %s export to %s.", exportFormat, exportFilename)
	collectedContent := rebuildForStringsExport(twittererList, concat)
	exportFile(collectedContent)
}

func ExportJson(exportFilename string, twittererList []string, exportFormat string) {
	fmt.Printf("Starting %s export to %s.", exportFormat, exportFilename)
	collectedContent := collectContent(twittererList)
	exportFile(string(collectedContent))
}

func ExportXml(exportFilename string, twittererList []string, exportFormat string) {
	fmt.Printf("Starting %s export to %s.", exportFormat, exportFilename)
	xmlContent, err := xml.Marshal(twittererList)
	helpers.Check(err)
	exportFile(string(xmlContent))
}

func collectContent(twittererList []string) []byte {
	collectedContent, err := json.Marshal(twittererList)
	helpers.Check(err)
	return collectedContent
}

func rebuildForStringsExport(twittererList []string, concat string) string {
	var collectedContent string
	for _, twitterAccount := range twittererList {
		collectedContent = twitterAccount + concat + collectedContent
	}
	collectedContent = strings.TrimRight(collectedContent, concat)
	return collectedContent
}

func printTwitterAccount(twittererList []twitter.User, twitterCollection string) {
	for i, twitterer := range twittererList {

		twitterCollection += strconv.Itoa(i) + ": " + twitterer.ScreenName + " (" + twitterer.Name + ")" +
			" Verified: " + strconv.FormatBool(twitterer.Verified) + " Language: " + twitterer.Lang + "\n" +
			"   Description: " + twitterer.Description + "\n"

		twitterCollection += "\n---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----\n"
	}
	fmt.Println(twitterCollection)
}
