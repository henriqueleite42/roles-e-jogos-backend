package event_delivery_http

import (
	"net/http"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/adapters"
	event_usecase "github.com/henriqueleite42/roles-e-jogos-backend/internal/usecase/event"
	"github.com/rs/zerolog"
)

type eventController struct {
	logger       *zerolog.Logger
	authAdapter  adapters.Auth
	validator    adapters.Validator
	idAdapter    adapters.Id
	eventUsecase event_usecase.EventUsecase
}

type AddEventControllerInput struct {
	Mux    *http.ServeMux
	Logger *zerolog.Logger

	AuthAdapter adapters.Auth
	Validator   adapters.Validator
	IdAdapter   adapters.Id

	EventUsecase event_usecase.EventUsecase
}

func AddEventController(i *AddEventControllerInput) {
	eventController := &eventController{
		logger:       i.Logger,
		authAdapter:  i.AuthAdapter,
		validator:    i.Validator,
		idAdapter:    i.IdAdapter,
		eventUsecase: i.EventUsecase,
	}

	i.Mux.HandleFunc("/event/attendance", eventController.EventAttendance)
	i.Mux.HandleFunc("/event/next", eventController.EventNext)
}
