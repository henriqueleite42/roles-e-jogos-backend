package s3

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/adapters"
	"github.com/rs/zerolog"
)

type s3Adapter struct {
	logger    *zerolog.Logger
	client    *s3.Client
	presigner *s3.PresignClient
}

func NewS3(logger *zerolog.Logger) (adapters.Storage, error) {
	newLogger := logger.With().Str("adapter", "S3Adapter").Logger()

	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		return nil, err
	}

	s3Client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = os.Getenv("ENV") != "prod"
	})

	presinger := s3.NewPresignClient(s3Client)

	return &s3Adapter{
		logger:    &newLogger,
		client:    s3Client,
		presigner: presinger,
	}, nil
}
