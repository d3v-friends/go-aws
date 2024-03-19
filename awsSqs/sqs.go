package awsSqs

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type Manager struct {
	Client *sqs.Client
}

type CreateSQSManagerArgs interface {
	GetAccess() string
	GetSecret() string
	GetRegion() string
}

func CreateSQSManager(i CreateSQSManagerArgs) (manager *Manager, err error) {
	var cfg = aws.Config{
		Credentials: credentials.NewStaticCredentialsProvider(
			i.GetAccess(),
			i.GetSecret(),
			"",
		),
		Region:       i.GetRegion(),
		DefaultsMode: aws.DefaultsModeStandard,
	}

	manager = &Manager{
		Client: sqs.NewFromConfig(cfg),
	}

	return
}

// NewQueue FIFO 큐 생성
func (x *Manager) NewQueue(
	ctx context.Context,
	queueNm string,
) (err error) {

}
