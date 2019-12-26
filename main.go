package main

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"os"
)

func main() {

	// Keys set as environment variables(for secruity puposes obviously)
	consumerKey := os.Getenv("CONSUMER_ACCESS")
	consumerSecret := os.Getenv("CONSUMER_SECRET")
	accessToken := os.Getenv("ACCESS_TOKEN")
	accessSecret := os.Getenv("ACCESS_SECRET")

	// Oauth authetication
	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	//get ID of user
	userParams := &twitter.UserShowParams{ScreenName: "uzumakithegod"}
	user, _, _ := client.Users.Show(userParams)

	// Set up to get access to tweets associated with User ID
	tweets, _, _ := client.Timelines.UserTimeline(&twitter.UserTimelineParams{
		UserID: user.ID,
	})

	// loop through all tweets
	for i := 0; i < len(tweets); i++ {

		// Delete the current tweets iteration
		client.Statuses.Destroy(tweets[i].ID, nil)
	}

}
