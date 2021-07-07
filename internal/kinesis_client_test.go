package internal

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kinesis"
	"os"
	"testing"
)

func init() {
	_ = os.Setenv("KINESIS_STREAM_NAME", "feed-mock")
	_ = os.Setenv("KINESIS_STREAM_REGION", "aws-region")
}

type KinesisMockAPI struct{}

func (mock KinesisMockAPI) PutRecord(ctx context.Context,
	input *kinesis.PutRecordInput,
	optFns ...func(*kinesis.Options)) (*kinesis.PutRecordOutput, error) {

	output := &kinesis.PutRecordOutput{
		ShardId: aws.String("shard-01"),
	}
	return output, nil
}

func TestClient_Send(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		ctx  context.Context
		api  KdsApi
		data interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "Should send message to kinesis data stream",
			fields: fields{NewClient()},
			args: args{
				ctx:  context.Background(),
				api:  &KinesisMockAPI{},
				data: nil,
			}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.fields.client
			if err := c.sendWithApi(tt.args.ctx, tt.args.api, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("SendWithApi() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
