package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

const (
	RRID       string = "rr_id"
	RRJob      string = "rr_job"
	RRHeaders  string = "rr_headers"
	RRPipeline string = "rr_pipeline"
	RRTimeout  string = "rr_timeout"
	RRDelay    string = "rr_delay"
	StringType string = "String"
)

type data2 struct {
	aaa string
	bbb string
}

type data struct {
	ab map[string]data2
}

func main() {
	a := &sqs.SendMessageInput{
		MessageBody: aws.String("asdf"),
		QueueUrl:    aws.String("adsf"),
		// precision lost
		MessageAttributes: map[string]types.MessageAttributeValue{
			RRJob: {DataType: aws.String(StringType), BinaryListValues: nil, BinaryValue: nil, StringListValues: nil, StringValue: aws.String("asdf")},
		},
	}

	fmt.Println(a)
}
