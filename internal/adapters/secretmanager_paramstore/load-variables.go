package secretmanager_paramstore

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/utils"
)

// AwsName: LocalName
var VARIABLES_NAMES = []string{
	"Port",
	"DatabaseUrl",
	"WebsiteUrl",
	"LudopediaClientId",
	"LudopediaRedirectUri",
	"GoogleClientId",
	"GoogleRedirectUri",
	"MediasS3BucketName",
	"MediasCloudfrontUrl",
	"EmailAddressSystemMessages",
	"NameSystemMessages",
	"EmailTemplateSignInOtp",
	"CollectionImportPersonalCollectionFromLudopediaQueueId",
}

func (self *secretmanagerParamstore) loadVariables(cfg aws.Config) error {
	self.logger.Debug().Any("variables", VARIABLES_NAMES).Msg("variables to load")

	ssmClient := ssm.NewFromConfig(cfg)

	prefix := fmt.Sprintf("%s-", os.Getenv("ENV"))

	namesMap := make(map[string]string, len(VARIABLES_NAMES))
	secretsNamesAws := make([]string, 0, len(VARIABLES_NAMES))
	for _, v := range VARIABLES_NAMES {
		name := prefix + pascalToKebab(v)
		namesMap[name] = v
		secretsNamesAws = append(secretsNamesAws, name)
	}

	chunks := utils.Chunkfy(secretsNamesAws, 10)
	self.logger.Debug().Any("chunks", chunks).Msg("variables in chunks")

	for _, chunk := range chunks {
		result, err := ssmClient.GetParameters(context.Background(), &ssm.GetParametersInput{
			Names: chunk,
		})
		if err != nil {
			self.logger.Error().Err(err).Msg("fail to get variables")
			return err
		}

		if len(result.InvalidParameters) > 0 {
			self.logger.Error().Any("invalidParameters", result.InvalidParameters).Msg("fail to load parameters")
			return fmt.Errorf("fail to load parameters: %s", result.InvalidParameters)
		}

		for _, v := range result.Parameters {
			if v.Name == nil {
				self.logger.Error().Any("v", v).Msg("fail to get variable, \"Name\" is nil")
				return errors.New("fail to get variable")
			}
			if v.Value == nil {
				self.logger.Error().Msgf("fail to get variable \"%s\", \"Value\" is nil", *v.Name)
				return fmt.Errorf("variable \"%s\" has no \"Value\"", *v.Name)
			}

			localName := namesMap[*v.Name]

			switch localName {
			case "Port":
				self.secrets.Port = *v.Value

			case "WebsiteUrl":
				self.secrets.WebsiteUrl = *v.Value

			case "DatabaseUrl":
				self.secrets.DatabaseUrl = *v.Value

			case "LudopediaClientId":
				self.secrets.LudopediaClientId = *v.Value
			case "LudopediaRedirectUri":
				self.secrets.LudopediaRedirectUri = *v.Value

			case "GoogleClientId":
				self.secrets.GoogleClientId = *v.Value
			case "GoogleRedirectUri":
				self.secrets.GoogleRedirectUri = *v.Value

			case "MediasS3BucketName":
				self.secrets.MediasS3BucketName = *v.Value

			case "MediasCloudfrontUrl":
				self.secrets.MediasCloudfrontUrl = *v.Value

			case "EmailAddressSystemMessages":
				self.secrets.EmailAddressSystemMessages = *v.Value
			case "NameSystemMessages":
				self.secrets.NameSystemMessages = *v.Value
			case "EmailTemplateSignInOtp":
				self.secrets.EmailTemplateSignInOtp = *v.Value

			case "CollectionImportPersonalCollectionFromLudopediaQueueId":
				self.secrets.CollectionImportPersonalCollectionFromLudopediaQueueId = *v.Value
			}
		}
	}

	self.logger.Trace().Msg("all variables loaded")

	return nil
}
