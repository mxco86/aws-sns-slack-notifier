package main

import (
	"fmt"
	"github.com/nlopes/slack"
)

// SlackPost - post a message to Slack
func SlackPost(token string, channel string, username string, msg string) error {
	// Create a Slack API client using a legacy token
	client := slack.New(token)

	block := slack.NewTextBlockObject(
		"mrkdwn",
		fmt.Sprintf("%s", msg),
		false, false)

	section := slack.NewSectionBlock(block, nil, nil)

	// Post the message to Slack as a user
	_, _, err := client.PostMessage(channel,
		slack.MsgOptionText(msg, false),
		slack.MsgOptionBlocks(section),
		slack.MsgOptionUsername(username),
		slack.MsgOptionIconEmoji(":man_dancing:"))

	if err != nil {
		return err
	}

	return nil
}
