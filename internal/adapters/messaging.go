package adapters

import (
	"context"
	"sync"
	"time"
)

// Start - Third Party

type S3ObjectCreatedEventRecordS3Bucket struct {
	Name string `json:"name"`
}

type S3ObjectCreatedEventRecordS3Object struct {
	Key  string `json:"key"`
	Size string `json:"size"`
}

type S3ObjectCreatedEventRecordS3 struct {
	Bucket *S3ObjectCreatedEventRecordS3Bucket `json:"bucket"`
	Object *S3ObjectCreatedEventRecordS3Object `json:"object"`
}

type S3ObjectCreatedEventRecord struct {
	EventTime         string                        `json:"eventTime"`
	EventName         string                        `json:"eventName"`
	RequestParameters any                           `json:"requestParameters"`
	S3                *S3ObjectCreatedEventRecordS3 `json:"s3"`
}

type S3ObjectCreatedEvent struct {
	Records []*S3ObjectCreatedEventRecord
}

// End - Third Party

type Event = any

type JsonEvent = []byte

type SendEventInput struct {
	ListenerId string
	EventName  string
	Event      Event
}

type CreateListenerInput struct {
	// Context to execute the listener
	Ctx context.Context
	// Waiting group to be sure that all the process ended before terminating
	Wg *sync.WaitGroup
	// The Queue Url / Arn / Identifier to be used
	ListenerId string
	// The interval to process the batch of events
	// This may not be respected if the MaxEvents is reached,
	// if the webhook receives MaxEvents before the Delay ends,
	// them the events will be processed
	// Default = 5 seconds
	Delay time.Duration
	// Timeout to process events, after this period the
	// processing will be considered failed and will return to the
	// queue to try to be processed again
	// Default = 10 seconds
	Timeout time.Duration
	// Maximum amount of events to process at the same time
	// This value affects how Delay works
	// Zero is not a valid value, ot must be between 1 and the default value
	// Default = 10 (Max for SQS Implementation)
	MaxEvents int
	// Function to process events
	Fn func(i []JsonEvent)
}

type Messaging interface {
	// Private events can be acceded only by the module that owns that queue (publishes SQS)
	SendPrivateEvent(i *SendEventInput) error

	// Public events can be accessed by all modules subscribed to that topic (publishes to SNS)
	SendPublicEvent(i *SendEventInput) error

	CreateListener(i *CreateListenerInput)
}
