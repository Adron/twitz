package twitTwitz

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/Adron/twitz/helpers"
	"io/ioutil"
	"strings"
)

func ExportXml(exportFilename string, twittererList []string, exportFormat string) {
	fmt.Printf("Starting xml export to %s.", exportFilename)
	xmlContent, err := xml.Marshal(twittererList)
	helpers.Check(err)
	header := xml.Header
	collectedContent := header + string(xmlContent)
	exportFile(collectedContent, exportFilename+"."+exportFormat)
}

func ExportCsv(exportFilename string, twittererList []string, exportFormat string) {
	fmt.Printf("Starting txt export to %s.", exportFilename)
	collectedContent := rebuildForExport(twittererList, ",")
	exportFile(collectedContent, exportFilename+"."+exportFormat)
}

func ExportTxt(exportFilename string, twittererList []string, exportFormat string) {
	fmt.Printf("Starting %s export to %s.", exportFormat, exportFilename)
	collectedContent := rebuildForExport(twittererList, "\n")
	exportFile(collectedContent, exportFilename+"."+exportFormat)
}

func ExportJson(exportFilename string, twittererList []string, exportFormat string) {
	fmt.Printf("Starting %s export to %s.", exportFormat, exportFilename)
	collectedContent := collectContent(twittererList)
	exportFile(string(collectedContent), exportFilename+"."+exportFormat)
}

func collectContent(twittererList []string) []byte {
	collectedContent, err := json.Marshal(twittererList)
	helpers.Check(err)
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
	helpers.Check(err)
}
