package sqs_sns

import (
	"context"
	"encoding/json"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/adapters"
)

func (self *sqsSnsImplementation) SendPublicEvent(i *adapters.SendEventInput) error {
	jsonData, err := json.Marshal(i.Event)
	if err != nil {
		return err
	}

	ctx, cf := context.WithTimeout(context.Background(), time.Second*5)
	defer cf()

	detail := string(jsonData)
	dataType := "String"
	_, err = self.snsClient.Publish(ctx, &sns.PublishInput{
		Message: &detail,
		MessageAttributes: map[string]types.MessageAttributeValue{
			"EventName": {
				DataType:    &dataType,
				StringValue: &i.EventName,
			},
		},
		TopicArn: &i.ListenerId,
	})
	if err != nil {
		return err
	}

	return nil
}
