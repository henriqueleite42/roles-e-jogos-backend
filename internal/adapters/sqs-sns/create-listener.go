package sqs_sns

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/adapters"
)

type SqsEvent struct {
	Message string
}

func (self *sqsSnsImplementation) CreateListener(i *adapters.CreateListenerInput) {
	maxEvents := 10 // Max for SQS
	if i.MaxEvents > 0 && i.MaxEvents < maxEvents {
		maxEvents = i.MaxEvents
	}

	delay := 5 * time.Second
	if i.Delay != delay && i.Delay > 0 {
		delay = i.Delay
	}

	self.logger.Debug().Dur("val", delay).Msg("delay")

	ch := make(chan *string)

	i.Wg.Add(1)
	go func() {
		defer i.Wg.Done()

		for {
			select {
			case <-i.Ctx.Done():
				self.logger.Debug().Msg("context done, stop polling")
				return
			default:
				self.logger.Trace().Msg("pull for messages")
				ctxPull, cancelPull := context.WithTimeout(i.Ctx, 21*time.Second)
				msgs, err := self.sqsClient.ReceiveMessage(ctxPull, &sqs.ReceiveMessageInput{
					QueueUrl:            &i.ListenerId,
					MaxNumberOfMessages: int32(maxEvents),
					WaitTimeSeconds:     20, // Max for SQS
					VisibilityTimeout:   int32(i.Timeout),
				})
				cancelPull()
				if err != nil {
					self.logger.Error().Err(err).Msg("error on fetching messages")
					return
				}

				if len(msgs.Messages) == 0 {
					self.logger.Trace().Msg("no messages found")
					continue
				}
				self.logger.Debug().
					Int("amount", len(msgs.Messages)).
					Msg("messages found")

				entries := []types.DeleteMessageBatchRequestEntry{}

				for _, v := range msgs.Messages {
					ch <- v.Body
					entries = append(entries, types.DeleteMessageBatchRequestEntry{
						Id:            v.MessageId,
						ReceiptHandle: v.ReceiptHandle,
					})
				}

				self.logger.Trace().Msg("deleting messages")
				ctxDel, cancelDel := context.WithTimeout(i.Ctx, 5*time.Second)
				output, err := self.sqsClient.DeleteMessageBatch(ctxDel, &sqs.DeleteMessageBatchInput{
					QueueUrl: &i.ListenerId,
					Entries:  entries,
				})
				cancelDel()
				if err != nil {
					self.logger.Error().Err(err).Msg("error on deleting messages")
					return
				}
				if len(output.Failed) > 0 {
					self.logger.Error().Err(err).Msg("error on deleting messages from queue")
					continue
				}
				self.logger.Trace().Msg("messages deleted")
			}
		}
	}()

	events := [][]byte{}
	// Creates a standby delay, so the ticker is
	// triggered only if it has received messages
	delayStandby := 720 * time.Hour
	isTickerActive := false
	ticker := time.NewTicker(delayStandby)
	startTicker := func() {
		if isTickerActive {
			return
		}

		isTickerActive = true
		ticker.Reset(delay)
	}
	stopTicker := func() {
		isTickerActive = false
		ticker.Reset(delayStandby)
	}
	batchProcess := func() {
		defer i.Wg.Done()
		stopTicker()

		self.logger.Trace().Msg("run batchProcess")

		if len(events) == 0 {
			return
		}

		self.logger.Debug().
			Str("listenerId", i.ListenerId).
			Int("len", len(events)).
			Any("events", events).
			Msg("processing events")

		localEvents := events
		events = [][]byte{}

		i.Fn(localEvents)
	}

	go func() {
		for event := range ch {
			// eventJsonDecoded := SqsEvent{}
			// self.logger.Debug().Str("event", *event).Msg("event")
			// err := json.Unmarshal([]byte(*event), &eventJsonDecoded)
			// if err != nil {
			// 	self.logger.Error().Err(err).Msg("error on unmarshalling messages")
			// 	continue
			// }

			// eventBytes := []byte(eventJsonDecoded.Message)
			// events = append(events, eventBytes)
			events = append(events, []byte(*event))
			startTicker()

			if len(events) >= i.MaxEvents {
				self.logger.Trace().Msg("max events reached")
				i.Wg.Add(1)
				batchProcess()
			}
		}
	}()

	for {
		select {
		case <-i.Ctx.Done():
			self.logger.Debug().Msg("cancel requested, processing latest messages")
			i.Wg.Add(1)
			batchProcess()
			return
		case <-ticker.C:
			self.logger.Trace().Msg("ticker triggered")
			i.Wg.Add(1)
			batchProcess()
		}
	}
}
