package main

import (
	"context"
	"os"
	"time"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/adapters/auth_postgres"
	go_validator "github.com/henriqueleite42/roles-e-jogos-backend/internal/adapters/go-validator"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/adapters/google"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/adapters/ludopedia"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/adapters/s3"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/adapters/secretmanager_paramstore"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/adapters/ses"
	sqs_sns "github.com/henriqueleite42/roles-e-jogos-backend/internal/adapters/sqs-sns"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/adapters/xid"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/delivery"
	http_delivery "github.com/henriqueleite42/roles-e-jogos-backend/internal/delivery/http"
	queue_delivery "github.com/henriqueleite42/roles-e-jogos-backend/internal/delivery/queue"
	account_repository "github.com/henriqueleite42/roles-e-jogos-backend/internal/repository/account"
	collection_repository "github.com/henriqueleite42/roles-e-jogos-backend/internal/repository/collection"
	game_repository "github.com/henriqueleite42/roles-e-jogos-backend/internal/repository/game"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/repository/queries"
	account_usecase "github.com/henriqueleite42/roles-e-jogos-backend/internal/usecase/account"
	collection_usecase "github.com/henriqueleite42/roles-e-jogos-backend/internal/usecase/collection"
	event_usecase "github.com/henriqueleite42/roles-e-jogos-backend/internal/usecase/event"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

func main() {
	// ----------------------------
	//
	// Logger
	//
	// ----------------------------

	// UNIX Time is faster and smaller than most timestamps
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	logLevel := os.Getenv("LOG_LEVEL")
	switch logLevel {
	case "TRACE":
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	case "DEBUG":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "INFO":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "WARN":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "ERROR":
		fallthrough
	default:
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	}

	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	logger.Trace().Msg("start")

	// ----------------------------
	//
	// Secrets
	//
	// ----------------------------

	logger.Trace().Msg("loading secrets")

	secretsAdapter, err := secretmanager_paramstore.
		NewSecretManagerParamStore(&logger)
	if err != nil {
		logger.Fatal().
			Err(err).
			Msg("fail to get secrets")
	}

	logger.Info().Msg("secrets loaded")

	// ----------------------------
	//
	// Databases
	//
	// ----------------------------

	logger.Trace().Msg("connecting to database")

	db, err := pgxpool.New(
		context.Background(),
		"postgres://"+secretsAdapter.DatabaseUsername+":"+secretsAdapter.DatabasePassword+"@"+secretsAdapter.DatabaseUrl+":5432/database",
	)
	if err != nil {
		logger.Fatal().
			Err(err).
			Msg("fail to connect to the database")
	}
	defer db.Close()

	sqlcQueries := queries.New(db)

	logger.Info().Msg("connected to database")

	// ----------------------------
	//
	// Adapters
	//
	// ----------------------------

	logger.Trace().Msg("initializing adapters")

	googleAdapter, err := google.NewGoogle(&logger, secretsAdapter)
	if err != nil {
		logger.Fatal().
			Err(err).
			Msg("fail to initialize GoogleAdapter")
	}
	ludopediaSignInAdapter, ludopediaAdapter, err := ludopedia.NewLudopedia(&logger, secretsAdapter)
	if err != nil {
		logger.Fatal().
			Err(err).
			Msg("fail to initialize LudopediaAdapter")
	}
	goValidatorAdapter, err := go_validator.NewGoValidator()
	if err != nil {
		logger.Fatal().
			Err(err).
			Msg("fail to initialize GoValidatorAdapter")
	}
	s3Adapter, err := s3.NewS3(&logger)
	if err != nil {
		logger.Fatal().
			Err(err).
			Msg("fail to initialize S3Adapter")
	}
	sesAdapter, err := ses.NewSes(&logger)
	if err != nil {
		logger.Fatal().
			Err(err).
			Msg("fail to initialize S3Adapter")
	}
	sqsSnsAdapter, err := sqs_sns.NewSqsSnsService(&sqs_sns.NewSqsSnsServiceInput{
		Logger: &logger,
	})
	if err != nil {
		logger.Fatal().
			Err(err).
			Msg("fail to initialize sqsSnsAdapter")
	}
	xidAdapter, err := xid.NewXid(&logger)
	if err != nil {
		logger.Fatal().
			Err(err).
			Msg("fail to initialize XidAdapter")
	}

	logger.Info().Msg("adapters initialized")

	// ----------------------------
	//
	// Repositories
	//
	// ----------------------------

	logger.Trace().Msg("initializing repositories")

	accountRepository, err := account_repository.NewAccountRepository(&account_repository.NewAccountRepositoryInput{
		Logger:         &logger,
		Queries:        sqlcQueries,
		IdAdapter:      xidAdapter,
		SecretsAdapter: secretsAdapter,
	})
	if err != nil {
		logger.Fatal().
			Err(err).
			Msg("fail to initialize AccountRepository")
	}
	collectionRepository, err := collection_repository.NewCollectionRepository(&collection_repository.NewCollectionRepositoryInput{
		Logger:         &logger,
		Queries:        sqlcQueries,
		SecretsAdapter: secretsAdapter,
	})
	if err != nil {
		logger.Fatal().
			Err(err).
			Msg("fail to initialize CollectionRepository")
	}
	gameRepository, err := game_repository.NewGameRepository(&game_repository.NewGameRepositoryInput{
		Logger:  &logger,
		Queries: sqlcQueries,
	})
	if err != nil {
		logger.Fatal().
			Err(err).
			Msg("fail to initialize GameRepository")
	}

	logger.Info().Msg("repositories initialized")

	// ----------------------------
	//
	// Services
	//
	// ----------------------------

	logger.Trace().Msg("initializing services")

	accountUsecase := &account_usecase.AccountUsecaseImplementation{
		Logger:            &logger,
		Db:                db,
		GoogleAdapter:     googleAdapter,
		LudopediaAdapter:  ludopediaSignInAdapter,
		IdAdapter:         xidAdapter,
		StorageAdapter:    s3Adapter,
		SecretsAdapter:    secretsAdapter,
		MessagingAdapter:  sqsSnsAdapter,
		EmailAdapter:      sesAdapter,
		AccountRepository: accountRepository,
	}
	collectionUsecase := &collection_usecase.CollectionUsecaseImplementation{
		Logger:               &logger,
		Db:                   db,
		LudopediaAdapter:     ludopediaAdapter,
		IdAdapter:            xidAdapter,
		StorageAdapter:       s3Adapter,
		SecretsAdapter:       secretsAdapter,
		EmailAdapter:         sesAdapter,
		MessagingAdapter:     sqsSnsAdapter,
		AccountRepository:    accountRepository,
		CollectionRepository: collectionRepository,
		GameRepository:       gameRepository,
	}
	eventUsecase := &event_usecase.EventUsecaseImplementation{
		Logger:         &logger,
		Db:             db,
		SecretsAdapter: secretsAdapter,
	}

	logger.Info().Msg("services initialized")

	// ----------------------------
	//
	// Deliveries
	//
	// ----------------------------

	logger.Trace().Msg("initializing deliveries")

	authPostgresAdapter, err := auth_postgres.NewAuthPostgres(&auth_postgres.NewAuthPostgresInput{
		Logger:            &logger,
		AccountRepository: accountRepository,
	})
	if err != nil {
		logger.Fatal().
			Err(err).
			Msg("fail to initialize AuthPostgresAdapter")
	}

	logger.Trace().Msg("initializing queue server")
	queueDelivery := queue_delivery.NewQueueDelivery(&queue_delivery.NewQueueDeliveryInput{
		Logger: &logger,

		IdAdapter:        xidAdapter,
		SecretsAdapter:   secretsAdapter,
		MessagingAdapter: sqsSnsAdapter,

		CollectionUsecase: collectionUsecase,
	})
	queueDelivery.Listen()
	logger.Info().Msg("queue server initialized")

	logger.Trace().Msg("initializing http server")
	httpDelivery := http_delivery.NewHttpDelivery(&http_delivery.NewHttpDeliveryInput{
		Logger: &logger,

		Validator:      goValidatorAdapter,
		SecretsAdapter: secretsAdapter,
		AuthAdapter:    authPostgresAdapter,
		IdAdapter:      xidAdapter,

		AccountUsecase:    accountUsecase,
		CollectionUsecase: collectionUsecase,
		EventUsecase:      eventUsecase,
	})
	httpDelivery.Listen()
	logger.Info().Msg("http server initialized")

	logger.Info().Msg("deliveries initialized")

	// ----------------------------
	//
	// Gracefully shutdown
	//
	// ----------------------------

	logger.Trace().Msg("setup gracefully shutdown")

	delivery.GracefullyShutdown(&delivery.GracefullyShutdownInput{
		Logger:  &logger,
		Timeout: 15 * time.Second,
		Deliveries: []delivery.Delivery{
			queueDelivery,
			httpDelivery,
		},
	})

	logger.Info().Msg("shutdown completed")
}
