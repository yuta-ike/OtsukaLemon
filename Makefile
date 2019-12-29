build:
	go get github.com/aws/aws-lambda-go/events
	go get github.com/aws/aws-lambda-go/lambda
	go get github.com/line/line-bot-sdk-go/linebot
	go build -o functions/webhook ./...