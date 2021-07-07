package internal

import (
	"context"
	"encoding/json"
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

func (c Client) Send(ctx context.Context, data interface{}) error {
	return c.SendWithApi(ctx, c.awsKinesis, data)
}

func (c Client) SendWithApi(ctx context.Context, api KdsApi, data interface{}) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if _, err = api.PutRecord(ctx, &kinesis.PutRecordInput{
		Data: bytes,
		// To use considering ordering if necessary PartitionKey: aws.String(data.Id),
		StreamName: aws.String(c.config.StreamName),
	}); err != nil {
		return err
	}

	return nil
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
