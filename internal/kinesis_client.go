package internal

type Config struct {
	StreamName string `env:"KINESIS_STREAM_NAME"`
	Region     string `env:"KINESIS_STREAM_REGION"`
}
