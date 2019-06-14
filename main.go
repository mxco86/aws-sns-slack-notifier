package main

import (
	"./notifier"
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"os"
)

type slackMessage struct {
	Token    string `json:"token"`
	Channel  string `json:"channel"`
	Username string `json:"username"`
}

// Lambda handler wrapper for the Slack notifier function. Context and Event
// are added by the execution environment
func handler(ctx context.Context, snsEvent events.SNSEvent) (err error) {

	// Slack-specific configuration is via Lambda environment variables
	slackMsg := slackMessage{
		Token:    os.Getenv("TOKEN"),
		Channel:  os.Getenv("CHANNEL"),
		Username: os.Getenv("USERNAME"),
	}

	// Event may contain multiple SNS records - send a message per-record
	for _, record := range snsEvent.Records {
		snsRecord := record.SNS

		token := slackMsg.Token
		channel := slackMsg.Channel
		msg := snsRecord.Message
		username := slackMsg.Username
		err = notifier.SlackPost(token, channel, username, msg)

		if err != nil {
			return err
		}
	}

	return
}

// Entrypoint for the lambda execution
func main() {
	lambda.Start(handler)
}
