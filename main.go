package main

import (
	"fmt"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func main() {

	consumerKey := os.Getenv("CONSUMER_ACCESS")
	consumerSecret := os.Getenv("CONSUMER_SECRET")
	accessToken := os.Getenv("ACCESS_TOKEN")
	accessSecret := os.Getenv("ACCESS_SECRET")

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	// // Twitter client
	client := twitter.NewClient(httpClient)

	//get ID of user
	userParams := &twitter.UserShowParams{ScreenName: "uzumakithegod"}
	user, _, _ := client.Users.Show(userParams)

	fmt.Println(user.ID)

	// //loop through all Tweets
	// tweet, _ := client.Timelines.UserTimeline(&twitter.UserTimelineParams{
	// 	UserID:
	// })

	//	delete all tweets
	//client.Statuses.Destroy(1209537877538988032, nil)

}
