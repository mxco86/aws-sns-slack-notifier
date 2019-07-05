package main

import (
	"encoding/json"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestFormatSlackMessage(t *testing.T) {

	// Create a test event
	sns := events.CloudWatchEvent{
		DetailType: "TestDetail",
		Detail:     json.RawMessage(`{"State": "TestState"}`),
	}

	msg, err := formatSlackMessage(sns)
	if err != nil {
		t.Errorf("Slack message failure")
	}

	// Two block element pairs should be built from the test event
	if len(msg) != 4 {
		t.Errorf("Slack message built incorrectly: %d", len(msg))
	}

	// Check the value text in the second block
	if msg[3].Text != "TestState" {
		t.Errorf("Slack message value error: %s", msg[3].Text)
	}
}
