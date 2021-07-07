# Go e Kinesis stream

## Implementação de um cliente de comunicação usando Go e Kinesis Stream (AWS)

- Kinesis Stream
- Golang

### Dependências

- AWS SDK Go v2
- Lib Variáves de Ambiente : github.com/caarlos0/env/v6


### Varáveis de Ambiente

	StreamName string `env:"KINESIS_STREAM_NAME"`
	Region     string `env:"KINESIS_STREAM_REGION"`