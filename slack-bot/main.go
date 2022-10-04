package main

import (
	"context"
	"fmt"
	"github.com/shomali11/slacker"
	"log"
	"os"
	"strconv"
)

func printCommandEvents(slackerEventChannel <-chan *slacker.CommandEvent) {
	for event := range slackerEventChannel {
		fmt.Println(event.Event)
		fmt.Println(event.Command)
	}
}

func main() {
	_ = os.Setenv("SLACK_BOT_TOKEN", "")
	_ = os.Setenv("SLACK_APP_TOKEN", "")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("my job is <year>", &slacker.CommandDefinition{
		Description: "",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				fmt.Println(err)
			}
			age := 2021 - yob
			r := fmt.Sprintf("age is %d", age)
			_ = response.Reply(r)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatalln(err)
	}
}
