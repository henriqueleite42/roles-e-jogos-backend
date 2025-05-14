package sqs_sns

import (
	"context"
	"encoding/json"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/adapters"
)

func (self *sqsSnsImplementation) SendPrivateEvent(i *adapters.SendEventInput) error {
	jsonData, err := json.Marshal(i.Event)
	if err != nil {
		return err
	}

	ctx, cf := context.WithTimeout(context.Background(), time.Second*5)
	defer cf()

	detail := string(jsonData)
	dataType := "String"
	_, err = self.sqsClient.SendMessage(ctx, &sqs.SendMessageInput{
		QueueUrl:    &i.ListenerId,
		MessageBody: &detail,
		MessageAttributes: map[string]types.MessageAttributeValue{
			"EventName": {
				DataType:    &dataType,
				StringValue: &i.EventName,
			},
		},
	})
	if err != nil {
		return err
	}

	return nil
}
