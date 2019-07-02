package main

import (
	"regexp"
	"testing"

	"github.com/mxco86/awstypes"
)

func TestFormatSlackMessage(t *testing.T) {

	sns := awstypes.SNSCodePipelineMessage{
		DetailType: "TestDetail",
	}

	sns.Detail.Pipeline = ""
	sns.Detail.Stage = ""
	sns.Detail.Action = "TestAction"
	sns.Detail.State = ""
	sns.Detail.ID = ""

	msg, err := formatSlackMessage(sns)
	if err != nil {
		t.Errorf("Slack message failure")
	}

	matched, err := regexp.MatchString(".*TestDetail\n.*TestAction", msg)
	if matched == false {
		t.Errorf("Slack message built incorrectly")
	}
}
