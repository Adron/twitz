package coreTwitz

import (
	"github.com/Adron/twitz/helpers"
	"github.com/spf13/viper"
	"io/ioutil"
	"regexp"
	"strings"
)

func BuildTwitterList(withAtSymbols bool) []string {
	theFile := viper.GetString("file")
	theTwitterers, err := ioutil.ReadFile(theFile)
	helpers.Check(err)
	stringTwitterers := string(theTwitterers[:])
	splitFields := strings.Fields(stringTwitterers)
	var completedTwittererList []string
	for _, aField := range splitFields {
		if strings.HasPrefix(aField, "@") && aField != "@" {
			reg, _ := regexp.Compile("[^a-zA-Z0-9_@]")
			processedString := reg.ReplaceAllString(aField, "")
			if withAtSymbols {
				processedString = strings.Trim(processedString, "@")
			}
			completedTwittererList = append(completedTwittererList, processedString)
		}
	}
	return completedTwittererList
}
