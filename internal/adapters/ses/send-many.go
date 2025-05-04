package ses

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ses"
	ses_types "github.com/aws/aws-sdk-go-v2/service/ses/types"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/adapters"
)

func (self *sesAdapter) SendMany(i *adapters.SendManyInput) error {
	self.logger.Trace().Msg("start SendMany")

	self.logger.Trace().Msg("try to build destinations")
	destinations := make([]ses_types.BulkEmailDestination, len(i.Contents))
	for k, v := range i.Contents {
		placeholdersBytes, err := json.Marshal(v.Placeholders)
		if err != nil {
			return err
		}
		placeholders := string(placeholdersBytes)

		destinations[k] = ses_types.BulkEmailDestination{
			Destination: &ses_types.Destination{
				ToAddresses: []string{v.RecipientAddress},
			},
			ReplacementTemplateData: &placeholders,
		}
	}
	self.logger.Debug().Interface("destinations", destinations).Msg("destinations built")

	source := fmt.Sprintf("%s <%s>", i.SenderName, i.SenderAddress)
	defaultTemplateData := "{}"
	input := &ses.SendBulkTemplatedEmailInput{
		Source:              &source,
		Template:            &i.TemplateId,
		Destinations:        destinations,
		DefaultTemplateData: &defaultTemplateData,
	}
	self.logger.Debug().Interface("input", input).Msg("input built")

	_, err := self.client.SendBulkTemplatedEmail(context.Background(), input)
	if err != nil {
		self.logger.Error().Err(err).Msg("fail to send emails")
		return err
	}

	self.logger.Trace().Msg("end SendMany")

	return nil
}
