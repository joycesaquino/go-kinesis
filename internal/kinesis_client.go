package internal

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kinesis"
	"github.com/caarlos0/env/v6"
	"log"
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

func NewClient() *Client {
	var config Config

	err := env.Parse(&config)
	if err != nil {
		log.Fatal("[ERROR] - Couldn't build settings for kinesis stream client", err)
	}

	return &Client{
		config:     config,
		awsKinesis: kinesis.NewFromConfig(aws.Config{Region: config.Region}),
	}
}
