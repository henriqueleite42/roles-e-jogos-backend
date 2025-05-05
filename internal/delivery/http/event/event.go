package event_delivery_http

import (
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/adapters"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/usecase/event"
	"github.com/rs/zerolog"
)

type eventController struct {
	logger        *zerolog.Logger
	validator     adapters.Validator
	eventUsecase  event_usecase.EventUsecase
}

type AddEventControllerInput struct {
	Logger        *zerolog.Logger
	Validator     adapters.Validator
	EventUsecase  event_usecase.EventUsecase
}

func AddEventController(i *AddEventControllerInput) {
	// Add routes here. Ex:
	// http.HandleFunc("/", self.Handler)
}
