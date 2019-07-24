package storage

import (
	"github.com/Adron/twitz/helpers"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/gocql/gocql"
	"github.com/satori/go.uuid"
	"github.com/spf13/viper"
)

func InsertTwitterAccount(cassieSession gocql.Session, twitterUser twitter.User) {
	newUUID := uuid.NewV4()
	preparedUUID, _ := gocql.ParseUUID(newUUID.String())

	helpers.Check(cassieSession.Query(`INSERT INTO twitz.twitteraccounts (id, username, name, createat, 
													description, email, followerscount, friendscount, following, twitterid, twitteridstr, listedcount, 
													location, geoEnabled, profileImageUri, profileBackgroundImageUri, statusesCount) 
													VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?) IF NOT EXISTS;`,
		preparedUUID,
		twitterUser.ScreenName,
		twitterUser.Name,
		twitterUser.CreatedAt,
		twitterUser.Description,
		twitterUser.Email,
		twitterUser.FollowersCount,
		twitterUser.FriendsCount,
		twitterUser.Following,
		twitterUser.ID,
		twitterUser.IDStr,
		twitterUser.ListedCount,
		twitterUser.Location,
		twitterUser.GeoEnabled,
		twitterUser.ProfileImageURL,
		twitterUser.ProfileBackgroundImageURL,
		twitterUser.StatusesCount).Exec())
}

func GetCassieSession() gocql.Session {
	cluster := gocql.NewCluster(viper.GetString("cassie"))
	cluster.Keyspace = viper.GetString("keyspace")
	cluster.Consistency = gocql.One

	session, _ := cluster.CreateSession()
	return *session
}

