package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
	"log"
	"os"
	"time"
)

func main() {
	godotenv.Load(".env")

	token := os.Getenv("SLACK_AUTH_TOKEN")
	channelID := os.Getenv("SLACK_CHANNEL_ID")

	fmt.Println("token:", token)
	fmt.Println("channelID:", channelID)

	api := slack.New(token)

	_, err := api.AddChannelReminder(channelID, "Have you send tour daily report?)", "at 15:00 every Monday, Tuesday, Wednesday, Thursday, Friday")

	if err != nil {
		fmt.Println("error in addchren")
		log.Fatalf("%s\n", err)
	}

	log.Printf("Message successfully sent to Channel %s at %s\n", channelID, time.ANSIC)
}
