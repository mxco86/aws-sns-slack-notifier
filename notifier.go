package main

import (
	"fmt"

	"github.com/nlopes/slack"

	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

func formatSlackMessage(inc events.CloudWatchEvent) (msg []*slack.TextBlockObject, err error) {

	// New slice to hold Slack Fields
	fieldSlice := make([]*slack.TextBlockObject, 0)

	// Unmarshal the event detail as it is imported as raw JSON
	var codePipelineEventDetail CodePipelineEventDetail
	JSONErr := json.Unmarshal([]byte(inc.Detail), &codePipelineEventDetail)
	if JSONErr != nil {
		return fieldSlice, fmt.Errorf("Input event error: %v", JSONErr)
	}

	fields := map[string]string{
		"Type":     inc.DetailType,
		"Action":   codePipelineEventDetail.Action,
		"Pipeline": codePipelineEventDetail.Pipeline,
		"State":    codePipelineEventDetail.State,
		"ID":       codePipelineEventDetail.ID,
	}

	for field, value := range fields {

		// Skip block creation if we have no value
		if value == "" {
			continue
		}

		// Create field and value blocks which will display in Slack as a table
		fieldBlock := slack.NewTextBlockObject(
			"mrkdwn",
			fmt.Sprintf("*%s*", field),
			false, false)
		valueBlock := slack.NewTextBlockObject(
			"plain_text",
			fmt.Sprintf("%s", value),
			false, false)

		fieldSlice = append(fieldSlice, fieldBlock)
		fieldSlice = append(fieldSlice, valueBlock)
	}

	return fieldSlice, nil
}

// SlackPost - post a message to Slack
func SlackPost(token string, channel string, username string, header string, fields []*slack.TextBlockObject) error {

	headerText := slack.NewTextBlockObject("mrkdwn", fmt.Sprintf("*%s*", header), false, false)
	headerSection := slack.NewSectionBlock(headerText, nil, nil)
	fieldsSection := slack.NewSectionBlock(nil, fields, nil)

	// Create a Slack API client using a legacy token
	client := slack.New(token)

	// Post the message to Slack as a user
	_, _, err := client.PostMessage(
		channel,
		slack.MsgOptionText(header, false),
		slack.MsgOptionBlocks(headerSection, fieldsSection),
		slack.MsgOptionUsername(username),
		slack.MsgOptionIconEmoji(":man_dancing:"))

	if err != nil {
		return err
	}

	return nil
}
