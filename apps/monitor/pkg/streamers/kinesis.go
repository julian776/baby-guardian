package streamers

import (
	"context"
	"errors"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kinesis"
	"github.com/aws/aws-sdk-go-v2/service/kinesis/types"
	awsInternal "github.com/julian776/baby-guardian/monitor/pkg/aws"
	"github.com/rs/zerolog"
)

var (
	errStreamIsBeingDeleted = errors.New("stream is being deleted")
)

type KinesisConfig struct {
	StreamName      string
	CredentialsFile string
	Region          string
	ShardCount      *int32
	StreamMode      types.StreamMode
}

type Kinesis struct {
	Config *KinesisConfig
	Client *kinesis.Client
	logger *zerolog.Logger
}

func NewKinesis(config *KinesisConfig, logger *zerolog.Logger) *Kinesis {
	c := aws.Config{
		Region:      config.Region,
		Credentials: awsInternal.GetCredentials(config.CredentialsFile),
	}

	return &Kinesis{
		Client: kinesis.NewFromConfig(c),
		Config: config,
		logger: logger,
	}
}

func (k *Kinesis) Start(ctx context.Context) error {
	k.logger.Info().Msg("initializing kinesis stream")

	err := k.createStream(ctx)
	if err != nil {
		return err
	}

	k.logger.Info().Msg("kinesis stream initialized")

	return nil
}

func (k *Kinesis) PutRecord(ctx context.Context, data []byte) error {
	_, err := k.Client.PutRecord(ctx, &kinesis.PutRecordInput{
		Data:         data,
		StreamName:   &k.Config.StreamName,
		PartitionKey: aws.String("1"),
	})
	if err != nil {
		return err
	}

	return nil
}

func (k *Kinesis) createStream(ctx context.Context) error {
	ok, err := k.validateStatus(ctx)
	if err != nil {
		return err
	}

	if ok {
		return nil
	}

	_, err = k.Client.CreateStream(ctx, &kinesis.CreateStreamInput{
		ShardCount: k.Config.ShardCount,
		StreamName: &k.Config.StreamName,
		StreamModeDetails: &types.StreamModeDetails{
			StreamMode: k.Config.StreamMode,
		},
	})
	if err != nil {
		return err
	}

	timeCtx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()
	if err := k.waitForStreamActive(timeCtx); err != nil {
		return err
	}

	return nil
}

func (k *Kinesis) waitForStreamActive(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		ok, err := k.validateStatus(ctx)
		if err != nil {
			return err
		}

		if ok {
			break
		}

		k.logger.Info().Msg("waiting for stream to become active")
		time.Sleep(1 * time.Second)
	}

	k.logger.Info().Msg("stream is active")
	return nil
}

func (k *Kinesis) validateStatus(ctx context.Context) (bool, error) {
	desc, err := k.Client.DescribeStream(ctx, &kinesis.DescribeStreamInput{
		StreamName: &k.Config.StreamName,
	})
	if err != nil {
		return false, err
	}

	switch desc.StreamDescription.StreamStatus {
	case types.StreamStatusActive:
		return true, nil
	case types.StreamStatusCreating:
		return false, nil
	case types.StreamStatusUpdating:
		return true, nil
	case types.StreamStatusDeleting:
		return false, errStreamIsBeingDeleted
	}

	return false, nil
}
