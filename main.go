package main

import (
	"./notifier"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"os"
)

type slackMessage struct {
	Token    string `json:"token"`
	Channel  string `json:"channel"`
	Username string `json:"username"`
}

type snsMessage struct {
	DetailType string `json:"detail-type"`
	Detail     struct {
		Pipeline string `json:"pipeline"`
		Stage    string `json:"stage"`
		Action   string `json:"action"`
		State    string `json:"state"`
		ID       string `json:"execution-id"`
	} `json:"detail"`
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

		// Process the raw JSON message
		slackFormattedMsg, _ := CreateSlackFormattedMsg(snsRecord.Message)

		token := slackMsg.Token
		channel := slackMsg.Channel
		msg := slackFormattedMsg
		username := slackMsg.Username
		err = notifier.SlackPost(token, channel, username, msg)

		if err != nil {
			return err
		}
	}

	return
}

// CreateSlackFormattedMsg -- a function
func CreateSlackFormattedMsg(SNSMsg string) (slackFormattedMsg string, err error) {

	var slackMessage snsMessage
	inputErr := json.Unmarshal([]byte(SNSMsg), &slackMessage)
	if inputErr != nil {
		fmt.Println(inputErr)
		return "", fmt.Errorf("Input event error: %v", inputErr)
	}

	ret, _ := json.Marshal(slackMessage)

	return string(ret), nil
}

// Entrypoint for the lambda execution
func main() {
	lambda.Start(handler)
}
