package main

import (
	// "fmt"
	"testing"
)

func TestSNSJSON(t *testing.T) {
	_, err := CreateSlackFormattedMsg(
		"{\"version\":\"0\",\"id\":\"c1093c59-c137-8e5f-df23-1a9d2a9fd46d\",\"detail-type\":\"CodePipeline Action Execution State Change\",\"source\":\"aws.codepipeline\",\"account\":\"068087527308\",\"time\":\"2019-06-26T16:46:15Z\",\"region\":\"eu-west-1\",\"resources\":[\"arn:aws:codepipeline:eu-west-1:068087527308:mxco86-web-site\"],\"detail\":{\"pipeline\":\"mxco86-web-site\",\"execution-id\":\"a6292101-a4e9-4516-86ea-cbaae5f65936\",\"stage\":\"ProdBuildAndDeploy\",\"action\":\"BuildAndDeploy\",\"state\":\"SUCCEEDED\",\"region\":\"eu-west-1\",\"type\":{\"owner\":\"AWS\",\"provider\":\"CodeBuild\",\"category\":\"Build\",\"version\":\"1\"},\"version\":2.0}}")

	if err != nil {
		t.Errorf("JSON Failure")
	}
}
