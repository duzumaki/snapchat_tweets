package main

import (
	"fmt"
	"log"
	"os"
	"strings"

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

	//time of tweet creation
	timeofTweet := tweets[1].CreatedAt
	//	fmt.Println(timeofTweet)

	//split strng
	timeofTweetSplitByWhiteSpace := strings.Fields(timeofTweet)

	dayofMonth := timeofTweetSplitByWhiteSpace[2]
	month := timeofTweetSplitByWhiteSpace[1]
	year := timeofTweetSplitByWhiteSpace[0]

	date := year + "-" + month + "-" + dayofMonth
	fmt.Println(date)

	// // loop through all tweets
	// for i := 0; i < len(tweets); i++ {
	// 	client.Statuses.Destroy(tweets[i].ID, nil)
	// }

	// }

	// func cronJobLondon() {
	// 	ldn, _ := time.LoadLocation("Europe/London")
	// 	c := cron.New(cron.WithLocation(ldn))
	// 	c.AddFunc("* * * * *", func() { fmt.Println("every minute test") })

	// }
}

func monthToDate(month string) string {
	switch month {
	case "Jan":
		return "01"
	case "Feb":
		return "02"
	case "Mar":
		return "03"
	case "Apr":
		return "04"
	case "May":
		return "05"
	case "June":
		return "06"
	case "Jul":
		return "07"
	case "Aug":
		return "08"
	case "Sep":
		return "09"
	case "Oct":
		return "10"
	case "Nov":
		return "11"
	case "Dec":
		return "12"
	}
	return ""
}
