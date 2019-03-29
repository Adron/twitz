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
	ExportCsvTxtWorkload(exportFilename, twittererList, exportFormat)
}

func ExportTxt(exportFilename string, twittererList []string, exportFormat string) {
	ExportCsvTxtWorkload(exportFilename, twittererList, exportFormat)
}

func ExportCsvTxtWorkload(exportFilename string, twittererList []string, exportFormat string) {
	fmt.Printf("Starting %s export to %s.", exportFormat, exportFilename)
	collectedContent := rebuildForStringsExport(twittererList, "\n")
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
	header := xml.Header
	collectedContent := header + string(xmlContent)
	exportFile(collectedContent)
}

func collectContent(twittererList []string) []byte {
	collectedContent, err := json.Marshal(twittererList)
	helpers.Check(err)
	return collectedContent
}

func rebuildForStringsExport(twittererList []string, concat string) string {
	var collectedContent string
	for _, twitterAccount := range twittererList {
		collectedContent = collectedContent + concat + twitterAccount
	}
	if concat == "," {
		collectedContent = strings.TrimLeft(collectedContent, concat)
	}
	return collectedContent
}

func printTwitterAccount(twittererList []twitter.User, twitterCollection string) {
	for i, twitterer := range twittererList {

		tweet := twitterer.Status

		twitterCollection += strconv.Itoa(i) + ": " + twitterer.ScreenName + " (" + twitterer.Name + ")" +
			" Verified: " + strconv.FormatBool(twitterer.Verified) + " Language: " + twitterer.Lang + "\n" +
			"   Description: " + twitterer.Description + "\n" +
			"  Counts:\n" +
			"   Followers: " + strconv.Itoa(twitterer.FollowersCount) +
			" Friends: " + strconv.Itoa(twitterer.FriendsCount) +
			" Favorites: " + strconv.Itoa(twitterer.FavouritesCount) +
			" Following: " + strconv.FormatBool(twitterer.Following) +
			" Tweets: " + strconv.Itoa(twitterer.StatusesCount) + "\n" +
			"   ID: " + strconv.FormatInt(twitterer.ID, 10) +
			" Listed: " + strconv.Itoa(twitterer.ListedCount) +
			" Follower Request Sent: " + strconv.FormatBool(twitterer.FollowRequestSent) +
			" Geo Enabled: " + strconv.FormatBool(twitterer.GeoEnabled) + "\n" +
			"   Location: " + twitterer.Location + "\n" +
			"   Time Zone: " + twitterer.Timezone + "\n" +
			"   User URL: " + twitterer.URL + "\n" +
			"   Profile URL: " + twitterer.ProfileImageURL + "\n" +
			"   Email: " + twitterer.Email

		twitterCollection += "\n\n"

		twitterCollection += "   Status Text: " + tweet.Text + "\n" +
			"     Full Text: " + tweet.FullText + "\n" +
			"   Language: " + tweet.Lang + "\n" +
			"   Created: " + tweet.CreatedAt + "\n" +
			"   Favorites: " + strconv.Itoa(tweet.FavoriteCount) +
			" Quoted: " + strconv.Itoa(tweet.QuoteCount) +
			" Retweets: " + strconv.Itoa(tweet.RetweetCount) +
			" Replies: " + strconv.Itoa(tweet.ReplyCount) + "\n" +
			" Source: " + tweet.Source + "\n"

		twitterCollection += "\n---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----\n"
	}
	fmt.Println(twitterCollection)
}
