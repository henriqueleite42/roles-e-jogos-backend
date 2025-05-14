package queue_delivery

import (
	"encoding/json"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/adapters"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/models"
)


func (self *queueDelivery) CollectionImportPersonalCollectionFromLudopedia() {
	go func() {
		self.messagingAdapter.CreateListener(&adapters.CreateListenerInput{
			ListenerId: self.secretsAdapter.CollectionImportPersonalCollectionFromLudopediaQueueId,
			Ctx:        self.ctx,
			Wg:         self.wg,
			Fn: func(i []adapters.JsonEvent) {
				var events []*models.ImportCollectionEvent
				for _, v := range i {
					event := &models.ImportCollectionEvent{}
					err := json.Unmarshal(v, &event)
					if err != nil {
						self.logger.Error().Err(err).Msg(err.Error())
						continue
					}

					events = append(events, event)
				}

				self.logger.Debug().
					Str("func", "ImportPersonalCollectionFromLudopedia").
					Any("events", events).
					Msg("processing events")

				self.collectionUsecase.ImportPersonalCollectionFromLudopedia(self.ctx, events)
			},
		})
	}()
}

