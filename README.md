# AWS SNS Slack Notifier

## What Does This Do?
A simple Lambda to receive messages from CodePipeline, format the details into
a readable Slack message and post the message to a Slack instance.

## Configuration

The Slack token, channel and posting username must be configured via
environment variables in the Lambda instance. Values for these variables can
be added to the CloudFormation stack via the SAM configuration file before
building

| Env Var  | Purpose                                                 |
|----------|---------------------------------------------------------|
| CHANNEL  | Internal id of the Slack channel to post the message to |
| USERNAME | Name under which the message will be posted             |
| TOKEN    | Legacy token with which to access the Slack API         |

## Build and Deployment
The Go binary must be built before deployment with the correct
compilation flags for Lambda execution

```sh
# Build Go binary for Lambda execution
GOOS=linux go build -ldflags="-s -w"
```

The Lambda stack is defined in a SAM configuration and can be built and
deployed using the standard SAM commands. The S3 bucket that holds the
deployment package can be any existing bucket.

```sh
# Stack creation
sam package --template-file sam.yaml --s3-bucket ${DeploymentBucketName} --output-template-file sam-pkg.yaml
sam deploy --template-file ./sam-pkg.yaml --stack-name ${StackName} --capabilities CAPABILITY_IAM

# Stack deletion
aws cloudformation delete-stack --stack-name ${StackName}
```

## Testing
The code can be tested locally using the standard SAM commands or directly
using a lambci docker container

```sh
# Run locally in docker container (after building and from the root directory)
docker run --rm -v "$PWD":/var/task \
       --env CHANNEL="A Slack Channel ID" \
       --env USERNAME="An Arbitrary Username to Post As" \
       --env TOKEN="A Slack Legacy API Token" \
       lambci/lambda:go1.x aws-sns-slack-notifier "$(< ./test_input)"
```
