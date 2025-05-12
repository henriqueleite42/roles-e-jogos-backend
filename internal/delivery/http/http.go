package http_delivery

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/rs/cors"
	"github.com/rs/zerolog"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/adapters"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/delivery"
	account_delivery_http "github.com/henriqueleite42/roles-e-jogos-backend/internal/delivery/http/account"
	collection_delivery_http "github.com/henriqueleite42/roles-e-jogos-backend/internal/delivery/http/collection"
	account_usecase "github.com/henriqueleite42/roles-e-jogos-backend/internal/usecase/account"
	collection_usecase "github.com/henriqueleite42/roles-e-jogos-backend/internal/usecase/collection"
	event_usecase "github.com/henriqueleite42/roles-e-jogos-backend/internal/usecase/event"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/utils"
)

type httpDelivery struct {
	mux    *http.ServeMux
	server *http.Server

	logger         *zerolog.Logger
	validator      adapters.Validator
	authAdapter    adapters.Auth
	secretsAdapter *adapters.Secrets
	idAdapter      adapters.Id

	accountUsecase    account_usecase.AccountUsecase
	collectionUsecase collection_usecase.CollectionUsecase
	eventUsecase      event_usecase.EventUsecase
}

type NewHttpDeliveryInput struct {
	Logger *zerolog.Logger

	Validator      adapters.Validator
	SecretsAdapter *adapters.Secrets
	AuthAdapter    adapters.Auth
	IdAdapter      adapters.Id

	AccountUsecase    account_usecase.AccountUsecase
	CollectionUsecase collection_usecase.CollectionUsecase
	EventUsecase      event_usecase.EventUsecase
}

func (self *httpDelivery) Name() string {
	return "HttpDelivery"
}

func (self *httpDelivery) Listen() {
	go func() {
		account_delivery_http.AddAccountController(&account_delivery_http.AddAccountControllerInput{
			Mux:            self.mux,
			Logger:         self.logger,
			Validator:      self.validator,
			AuthAdapter:    self.authAdapter,
			SecretsAdapter: self.secretsAdapter,
			IdAdapter:      self.idAdapter,
			AccountUsecase: self.accountUsecase,
		})
		collection_delivery_http.AddCollectionController(&collection_delivery_http.AddCollectionControllerInput{
			Mux:               self.mux,
			Logger:            self.logger,
			Validator:         self.validator,
			AuthAdapter:       self.authAdapter,
			SecretsAdapter:    self.secretsAdapter,
			IdAdapter:         self.idAdapter,
			CollectionUsecase: self.collectionUsecase,
		})

		self.logger.Info().
			Msgf("HTTP server initialized at %v", self.server.Addr)

		if err := self.server.ListenAndServe(); err != http.ErrServerClosed {
			self.logger.Error().
				Err(err).
				Msgf("error starting HTTP server")
			return
		}

		self.logger.Info().
			Msg("HTTP server stopped serving new connections")
	}()
}

func (self *httpDelivery) Cancel(timeout time.Duration) {
	// Create a context with a timeout for the shutdown process
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		if err := self.server.Shutdown(ctx); err != nil {
			self.logger.Error().
				Err(err).
				Msg("HTTP server shutdown error")
		}
	}()

	if err := utils.WaitWithTimeout(&wg, timeout); err != nil {
		self.logger.Error().Err(err).Msg("http delivery shutdown timeout")
	}
}

func NewHttpDelivery(i *NewHttpDeliveryInput) delivery.Delivery {
	port := fmt.Sprintf(":%v", i.SecretsAdapter.Port)

	mux := http.NewServeMux()

	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{i.SecretsAdapter.WebsiteUrl}, // or ["*"] for all
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "Access-Control-Allow-Credentials"},
		AllowCredentials: true, // only if using cookies
	}).Handler(mux)

	server := &http.Server{
		Addr:    port,
		Handler: handler,
		// Optionally configure timeouts for graceful shutdown
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &httpDelivery{
		mux:               mux,
		server:            server,
		logger:            i.Logger,
		validator:         i.Validator,
		authAdapter:       i.AuthAdapter,
		secretsAdapter:    i.SecretsAdapter,
		idAdapter:         i.IdAdapter,
		accountUsecase:    i.AccountUsecase,
		collectionUsecase: i.CollectionUsecase,
		eventUsecase:      i.EventUsecase,
	}
}
