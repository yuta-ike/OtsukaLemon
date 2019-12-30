package main

import (
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/line/line-bot-sdk-go/linebot"
)

func handler(awsRequest events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	accessToken := os.Getenv("LINEBOT_ACCESS_TOKEN")
	secret := os.Getenv("LINEBOT_CHANNEL_SECRET")

	bot, err := linebot.New(secret, accessToken)
	if err != nil {
		log.Fatal(err)
	}

	event, err := UnmarshalLineRequest([]byte(awsRequest.Body))
	if err != nil {
		log.Fatal(err)
	}

	for _, event := range event.Events {
		replyToken := event.ReplyToken

		var message string
		if event.Message.Type == "text" {
			message = buildReply(event.Message.Text)
		} else {
			message = "テキストメッセージを送ってね"
		}

		if message == "おつかれもん" {
			if _, err = bot.ReplyMessage(replyToken, linebot.NewTextMessage(message), linebot.NewImageMessage("https://4.bp.blogspot.com/-ICHHirmVkJQ/Vq89AnGcG_I/AAAAAAAA3kA/iiDZGFKOEiE/s800/fruit_slice09_lemon.png", "https://4.bp.blogspot.com/-ICHHirmVkJQ/Vq89AnGcG_I/AAAAAAAA3kA/iiDZGFKOEiE/s800/fruit_slice09_lemon.png")).Do(); err != nil {
				log.Fatal(err)
			}
		} else {
			if _, err = bot.ReplyMessage(replyToken, linebot.NewTextMessage(message)).Do(); err != nil {
				log.Fatal(err)
			}
		}

	}

	log.Printf("handler3")

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       awsRequest.Body,
	}, nil
}

func main() {
	log.Printf("main func start")
	lambda.Start(handler)
}
