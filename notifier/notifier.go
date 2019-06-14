// Package notifier provides method to post a message to a specified Slack channel
package notifier

import (
	"github.com/nlopes/slack"
)

// SlackPost - post a message to Slack
func SlackPost(token string, channel string, username string, msg string) error {
	// Create a Slack API client using a legacy token
	client := slack.New(token)

	// Post the message to Slack as a user
	_, _, err := client.PostMessage(channel,
		slack.MsgOptionText(msg, false),
		slack.MsgOptionUsername(username),
		slack.MsgOptionIconEmoji(":man_dancing:"))

	if err != nil {
		return err
	}

	return nil
}
