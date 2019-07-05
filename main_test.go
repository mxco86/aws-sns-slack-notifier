package main

import (
	"encoding/json"
	// "regexp"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestFormatSlackMessage(t *testing.T) {

	sns := events.CloudWatchEvent{
		DetailType: "TestDetail",
		Detail:     json.RawMessage(`{"Stage": "TestStage"}`),
	}

	msg, err := formatSlackMessage(sns)
	if err != nil {
		t.Errorf("Slack message failure")
	}

	if len(msg) != 2 {
		t.Errorf("Slack message built incorrectly: %d", len(msg))
	}

	// matched, err := regexp.MatchString(".*TestDetail\n.*\n.*TestStage", msg)
	// if matched == false {
	// 	t.Errorf("Slack message built incorrectly")
	// }
}
