package event_delivery_http

import (
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/adapters"
	event_usecase "github.com/henriqueleite42/roles-e-jogos-backend/internal/usecase/event"
	"github.com/rs/zerolog"
)

type eventController struct {
	logger       *zerolog.Logger
	validator    adapters.Validator
	idAdapter    adapters.Id
	eventUsecase event_usecase.EventUsecase
}

type AddEventControllerInput struct {
	Logger *zerolog.Logger

	Validator adapters.Validator
	IdAdapter adapters.Id

	EventUsecase event_usecase.EventUsecase
}

func AddEventController(i *AddEventControllerInput) {
	eventController := &eventController{
		logger:       i.Logger,
		validator:    i.Validator,
		idAdapter:    i.IdAdapter,
		eventUsecase: i.EventUsecase,
	}

	// Add routes here. Ex:
	// http.HandleFunc("/", self.Handler)
}
