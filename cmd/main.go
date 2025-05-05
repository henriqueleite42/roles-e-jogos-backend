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
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/adapters/xid"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/delivery"
	http_delivery "github.com/henriqueleite42/roles-e-jogos-backend/internal/delivery/http"
	account_repository "github.com/henriqueleite42/roles-e-jogos-backend/internal/repository/account"
	collection_repository "github.com/henriqueleite42/roles-e-jogos-backend/internal/repository/collection"
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

	db, err := pgxpool.New(context.Background(), secretsAdapter.DatabaseUrl)
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
	ludopediaAdapter, err := ludopedia.NewLudopedia(&logger, secretsAdapter)
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
		Logger:  &logger,
		Queries: sqlcQueries,
	})
	if err != nil {
		logger.Fatal().
			Err(err).
			Msg("fail to initialize AccountRepository")
	}
	collectionRepository, err := collection_repository.NewCollectionRepository(&collection_repository.NewCollectionRepositoryInput{
		Logger:  &logger,
		Queries: sqlcQueries,
	})
	if err != nil {
		logger.Fatal().
			Err(err).
			Msg("fail to initialize CollectionRepository")
	}

	logger.Info().Msg("repositories initialized")

	// ----------------------------
	//
	// Adapters (repository dependents)
	//
	// ----------------------------

	authPostgresAdapter, err := auth_postgres.NewAuthPostgres(&auth_postgres.NewAuthPostgresInput{
		Logger:            &logger,
		AccountRepository: accountRepository,
	})
	if err != nil {
		logger.Fatal().
			Err(err).
			Msg("fail to initialize AuthPostgresAdapter")
	}

	// ----------------------------
	//
	// Services
	//
	// ----------------------------

	logger.Trace().Msg("initializing services")

	accountUsecase := &account_usecase.AccountUsecaseImplementation{
		Logger:            &logger,
		Db:                db,
		AccountRepository: accountRepository,
		GoogleAdapter:     googleAdapter,
		IdAdapter:         xidAdapter,
		StorageAdapter:    s3Adapter,
		SecretsAdapter:    secretsAdapter,
	}
	collectionUsecase := &collection_usecase.CollectionUsecaseImplementation{
		Logger:               &logger,
		Db:                   db,
		CollectionRepository: collectionRepository,
		LudopediaAdapter:     ludopediaAdapter,
		IdAdapter:            xidAdapter,
		StorageAdapter:       s3Adapter,
		SecretsAdapter:       secretsAdapter,
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

	logger.Trace().Msg("initializing http server")
	httpDelivery := http_delivery.NewHttpDelivery(&http_delivery.NewHttpDeliveryInput{
		Logger: &logger,

		Validator:      goValidatorAdapter,
		SecretsAdapter: secretsAdapter,
		AuthAdapter:    authPostgresAdapter,

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
			httpDelivery,
		},
	})

	logger.Info().Msg("shutdown completed")
}
