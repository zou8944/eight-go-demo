package main

import (
	"fmt"
	"github.com/slack-go/slack"
	"os"
)

func main() {
	_ = os.Setenv("SLACK_BOT_TOKEN", "")
	_ = os.Setenv("CHANNEL_ID", "")

	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"ZIPL.pdf"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			File:     fileArr[i],
			Channels: channelArr,
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Name: %s, URL: %s\n", file.Name, file.URL)
	}
}
