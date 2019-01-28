package twitTwitz

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
)

func PrintUserToConsole(twitterUser twitter.User) {
	fmt.Printf("Screenname: %s  Name: %s\n", twitterUser.ScreenName, twitterUser.Name)
	fmt.Printf("Followers: %d  Following: %d\n",
		twitterUser.FollowersCount,
		twitterUser.FriendsCount)
	fmt.Println("...")
}

func PrintUsersToConsole(twitterUsers []twitter.User) {
	for _, twitterUser := range twitterUsers {
		PrintUserToConsole(twitterUser)
	}
	fmt.Println("... ... ... done.")
}
