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
	// userTweetParams := &twitter.UserTimelineParams{UserID: user.ID}
	// tweets, _, err := client.Timelines.UserTimeline(userTweetParams)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	likedTweetParams := &twitter.FavoriteListParams{UserID: user.ID, Count: 200}
	likedTweets, _, _ := client.Favorites.List(likedTweetParams)

	for i := 0; i < len(likedTweets); i++ {
			destroyParms := &twitter.FavoriteDestroyParams{ID: likedTweets[i].ID}
			client.Favorites.Destroy(destroyParms)	
	}

	// go deleteTweets(tweets, client)
	// deleteLikes(tweets, client, user)
}

// func deleteTweets(tweets []twitter.Tweet, client *twitter.Client){

// 	timeNow := time.Now()

// 	for i := 0; i < len(tweets); i++ {
// 		go func(i int) {
// 			timeofTweet, err := tweets[i].CreatedAtTime()
// 			if err != nil{
// 				log.Fatal(err)
// 			}

// 	// get difference of time now and the time the tweet was created
// 			difference := timeNow.Sub(timeofTweet)
// 			if difference > time.Hour*24 {
// 				client.Statuses.Destroy(tweets[i].ID, nil)
// 				}
// 		}(i)
// }
// }

// func deleteLikes(tweets []twitter.Tweet, client *twitter.Client, user *twitter.User){

// }
