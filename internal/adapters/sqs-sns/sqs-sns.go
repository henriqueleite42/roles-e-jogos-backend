package sqs_sns

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/adapters"
	"github.com/rs/zerolog"
)

type sqsSnsImplementation struct {
	logger *zerolog.Logger

	sqsClient *sqs.Client
	snsClient *sns.Client
}

type NewSqsSnsServiceInput struct {
	Logger *zerolog.Logger
}

func NewSqsSnsService(i *NewSqsSnsServiceInput) (adapters.Messaging, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}

	sqsClient := sqs.NewFromConfig(cfg)
	snsClient := sns.NewFromConfig(cfg)

	return &sqsSnsImplementation{
		logger:    i.Logger,
		sqsClient: sqsClient,
		snsClient: snsClient,
	}, nil
}
