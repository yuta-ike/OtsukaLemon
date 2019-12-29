package main

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/line/line-bot-sdk-go/linebot"
)

func handler(awsRequest events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	accessToken := "52gP1+MnOGvHWEpjzixawGUIW1icZTs3ULhetUCjv6KIad4WZsrLw5whm2OJ62CmQkEIy3xnYXytBcK4jSR1chcZvuXGdWLxEcruW/ZZ5qYs/lGz3JUCbUKZmKnVQU6ecX97O4OXev+QrBbA9AFWkQdB04t89/1O/w1cDnyilFU="
	secret := "1aac82b5d4e864f29ab5008997bbb9be"
	bot, err := linebot.New(secret, accessToken)
	if err != nil {
		// log.Fatal(err)
	}

	event, err := UnmarshalLineRequest([]byte(awsRequest.Body))
	if err != nil {
		log.Fatal(err)
	}

	for _, event := range event.Event {
		replyToken := event.ReplyToken
		bot.ReplyMessage(replyToken, linebot.NewTextMessage("hello")).Do()
	}

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       awsRequest.Body,
	}, nil
}

func main() {
	lambda.Start(handler)
}
