package streamers

import (
	"context"
	"errors"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kinesis"
	"github.com/aws/aws-sdk-go-v2/service/kinesis/types"
	iaws "github.com/julian776/baby-guardian/libs/aws"
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
	Config            *KinesisConfig
	Client            *kinesis.Client
	StreamArn         string
	NextShardIterator *string
	ShardID           *string
	logger            *zerolog.Logger
}

func NewKinesis(config *KinesisConfig, logger *zerolog.Logger) *Kinesis {
	c := aws.Config{
		Region:      config.Region,
		Credentials: iaws.GetCredentials(config.CredentialsFile),
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
		Data:            data,
		StreamName:      &k.Config.StreamName,
		StreamARN:       &k.StreamArn,
		PartitionKey:    aws.String("1"),
		ExplicitHashKey: aws.String("1"),
	})
	if err != nil {
		return err
	}

	return nil
}

func (k *Kinesis) GetRecords(ctx context.Context) ([][]byte, error) {
	if k.NextShardIterator == nil {
		res, err := k.Client.GetShardIterator(ctx, &kinesis.GetShardIteratorInput{
			ShardId:           k.ShardID,
			ShardIteratorType: types.ShardIteratorTypeLatest,
			StreamName:        &k.Config.StreamName,
			StreamARN:         &k.StreamArn,
		})
		if err != nil {
			return make([][]byte, 0), err
		}

		k.NextShardIterator = res.ShardIterator
	}

	res, err := k.Client.GetRecords(ctx, &kinesis.GetRecordsInput{
		ShardIterator: k.NextShardIterator,
	})
	if err != nil {
		return make([][]byte, 0), err
	}

	k.NextShardIterator = res.NextShardIterator

	d := make([][]byte, len(res.Records))
	for i, r := range res.Records {
		d[i] = r.Data
	}

	return d, nil
}

func (k *Kinesis) createStream(ctx context.Context) error {
	exists, err := k.validateSreamExists(ctx)
	if err != nil {
		return err
	}

	if !exists {
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

		k.logger.Debug().Msg("waiting for stream to become active")
		time.Sleep(1 * time.Second)
	}

	return nil
}

func (k *Kinesis) validateStatus(ctx context.Context) (bool, error) {
	desc, err := k.Client.DescribeStream(ctx, &kinesis.DescribeStreamInput{
		StreamName: &k.Config.StreamName,
	})
	if err != nil {
		return false, err
	}

	k.ShardID = desc.StreamDescription.Shards[0].ShardId
	k.StreamArn = *desc.StreamDescription.StreamARN

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

func (k *Kinesis) validateSreamExists(ctx context.Context) (bool, error) {
	list, err := k.Client.ListStreams(ctx, &kinesis.ListStreamsInput{})
	if err != nil {
		return false, err
	}

	for _, stream := range list.StreamNames {
		if stream == k.Config.StreamName {
			return true, nil
		}
	}

	return false, nil
}
