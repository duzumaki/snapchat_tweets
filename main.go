package main

import (
	"log"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
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
	user, _, err := client.Users.Show(userParams)
	if err != nil {
		log.Fatal(err)
	}

	// Set up to get access to tweets associated with User ID
	userTweetParams := &twitter.UserTimelineParams{UserID: user.ID}
	tweets, _, err := client.Timelines.UserTimeline(userTweetParams)
	if err != nil {
		log.Fatal(err)
	}

	// loop through all tweets
	for i := 0; i < len(tweets); i++ {

		// Delete the current tweets iteration
		client.Statuses.Destroy(tweets[i].ID, nil)
	}

}
