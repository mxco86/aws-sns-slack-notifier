package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// SlackChannel -- A channel
type SlackChannel struct {
	Token    string `json:"token"`
	Channel  string `json:"channel"`
	Username string `json:"username"`
}

// CodePipelineEventDetail -- A codepipeline detail
type CodePipelineEventDetail struct {
	Pipeline string `json:"pipeline"`
	Stage    string `json:"stage"`
	Action   string `json:"action"`
	State    string `json:"state"`
	ID       string `json:"execution-id"`
}

// Lambda handler wrapper for the Slack notifier function. Context and Event
// are added by the execution environment
func handler(ctx context.Context, snsEvent events.SNSEvent) (err error) {

	// Slack-specific configuration is via Lambda environment variables
	slackChannel := SlackChannel{
		Token:    os.Getenv("TOKEN"),
		Channel:  os.Getenv("CHANNEL"),
		Username: os.Getenv("USERNAME"),
	}

	// Event may contain multiple SNS records - send a message per-record
	for _, record := range snsEvent.Records {
		snsRecord := record.SNS

		// Process the raw incoming JSON message into a struct
		var snsCodePipelineMessage events.CloudWatchEvent
		inputErr := json.Unmarshal([]byte(snsRecord.Message), &snsCodePipelineMessage)
		if inputErr != nil {
			return fmt.Errorf("Input event error: %v", inputErr)
		}

		// Format the slack message
		msg, _ := formatSlackMessage(snsCodePipelineMessage)

		// Send the slack message to the configured channel
		err = SlackPost(
			slackChannel.Token,
			slackChannel.Channel,
			slackChannel.Username,
			msg)

		if err != nil {
			return err
		}
	}

	return
}

func formatSlackMessage(inc events.CloudWatchEvent) (msg string, err error) {

	// Unmarshal the event detail as it is imported as raw JSON
	var codePipelineEventDetail CodePipelineEventDetail
	JSONErr := json.Unmarshal([]byte(inc.Detail), &codePipelineEventDetail)
	if JSONErr != nil {
		return "", fmt.Errorf("Input event error: %v", JSONErr)
	}

	return fmt.Sprintf(
		"*Type:* %s\n *Action:* %s",
		inc.DetailType,
		codePipelineEventDetail.Action), nil

}

// Entrypoint for the lambda execution
func main() {
	lambda.Start(handler)
}
