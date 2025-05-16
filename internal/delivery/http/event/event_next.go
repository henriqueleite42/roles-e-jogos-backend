package event_delivery_http

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/delivery/http/utils"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/models"
	event_usecase "github.com/henriqueleite42/roles-e-jogos-backend/internal/usecase/event"
)

func (self *eventController) EventNext(w http.ResponseWriter, r *http.Request) {
	reqId := self.idAdapter.GenReqId()

	logger := self.logger.With().
		Str("dmn", "Event").
		Str("mtd", "EventNext").
		Str("reqId", reqId).
		Logger()

	if r.Method == http.MethodGet {
		getNextEventsInput := &event_usecase.GetNextEventsInput{
			Pagination: models.GetDefaultPaginationInputTimestamp(),
		}

		query := r.URL.Query()

		afterQuery := query.Get("after")
		if afterQuery != "" {
			afterTime, err := time.Parse(time.RFC3339, afterQuery)
			if err == nil {
				getNextEventsInput.Pagination.After = &afterTime
			} else {
				logger.Info().Err(err).Msg("fail to convert after")
			}
		}
		limitQuery := query.Get("limit")
		if limitQuery != "" {
			limitInt, err := strconv.Atoi(limitQuery)
			if err == nil {
				getNextEventsInput.Pagination.Limit = limitInt
			} else {
				logger.Info().Err(err).Msg("fail to convert limit")
			}
		}

		logger.Trace().Msg("validate getNextEventsInput")
		err := self.validator.Validate(getNextEventsInput)
		if err != nil {
			logger.Info().Err(err).Msg("invalid getNextEventsInput")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		logger.Trace().Msg("create reqCtx")
		reqCtx := context.WithValue(context.Background(), "logger", logger)

		logger.Trace().Msg("call usecase")
		getNextEventsOutput, err := self.eventUsecase.GetNextEvents(reqCtx, getNextEventsInput)
		if err != nil {
			// If there are any errors that should be handled, add them here
			logger.Warn().Err(err).Msg("usecase err")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		logger.Trace().Msg("send response")
		utils.ZipAndSendResponse(&logger, w, getNextEventsOutput)
		logger.Trace().Msg("finish")
		return
	}

	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
}
