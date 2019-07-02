package main

import (
	"encoding/json"
	"regexp"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestFormatSlackMessage(t *testing.T) {

	sns := events.CloudWatchEvent{
		DetailType: "TestDetail",
		Detail:     json.RawMessage(`{"Action": "TestAction"}`),
	}

	msg, err := formatSlackMessage(sns)
	if err != nil {
		t.Errorf("Slack message failure")
	}

	matched, err := regexp.MatchString(".*TestDetail\n.*TestAction", msg)
	if matched == false {
		t.Errorf("Slack message built incorrectly")
	}
}
