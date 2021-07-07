package internal

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/kinesis"
)

type Config struct {
	StreamName string `env:"KINESIS_STREAM_NAME"`
	Region     string `env:"KINESIS_STREAM_REGION"`
}

type KdsApi interface {
	PutRecord(ctx context.Context, input *kinesis.PutRecordInput, optFns ...func(*kinesis.Options)) (*kinesis.PutRecordOutput, error)
}

type Client struct {
	config     Config
	awsKinesis KdsApi
}
