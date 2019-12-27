### Current Status: Working

Checks time of tweet and deletes it once it reaches 5 days in age (can be changed)

## Setup:

I run this on an Amazon EC2 instance and have a cron job running saying that every 5th day, to run the script. 
Since Go compiles the code when you run "go build x.go" you can point the cron job to run "x" directly.

You will need to have a twitter application setup in order to generate the consumer and access keys. Twitter 
usually approves applications fast: https://developer.twitter.com

```
e.g "crontab -e"
CONSUMER_ACCESS=xxx
CONSUMER_SECRET=xxx
ACCESS_TOKEN=xxx
ACCESS_SECRET=xx
0 0 */5 * * /usr/bin/main
```

Where "main" was "main.go". I ran "go build [folder_it's_in]/main.go to output the "main" script


I'll add documentation to this at some point, I'm too lazy. I'll reply to help if you tweet me @uzumakithegod
