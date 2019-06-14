# SNS Slack Notifier

```sh

# Build Go binary for Lambda execution
GOOS=linux go build -ldflags="-s -w"

# Run locally in docker container
docker run --rm -v "$PWD":/var/task \
--env CHANNEL="A Slack Channel ID" \
--env USERNAME="An Arbitrary Username to Post As" \
--env TOKEN="A Slack Legacy API Token" \
lambci/lambda:go1.x sns-slack-notifier \
"{ \"Records\": [ { \"Sns\": {\"Message\": \"A Test Message\"} },
 { \"Sns\": {\"Message\": \"Another Test Message\"} } ] }"

# Stack creation and deletion
sam package --template-file sam.yaml --s3-bucket ${DeploymentBucketName} --output-template-file sam-pkg.yaml
sam deploy --template-file ./sam-pkg.yaml --stack-name ${StackName} --capabilities CAPABILITY_IAM
aws cloudformation delete-stack --stack-name ${StackName}

```
