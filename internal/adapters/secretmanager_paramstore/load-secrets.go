package secretmanager_paramstore

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/utils"
)

// AwsName: LocalName
var SECRETS_NAMES = []string{
	"DatabaseUrl",
	"GoogleClientSecret",
}

func (self *secretmanagerParamstore) loadSecrets(cfg aws.Config) error {
	self.logger.Debug().Any("secrets", SECRETS_NAMES).Msg("secrets to load")

	secretManagerClient := secretsmanager.NewFromConfig(cfg)

	prefix := fmt.Sprintf("%s-", os.Getenv("ENV"))

	namesMap := make(map[string]string, len(SECRETS_NAMES))
	secretsNamesAws := make([]string, 0, len(SECRETS_NAMES))
	for _, v := range SECRETS_NAMES {
		name := prefix + pascalToKebab(v)
		namesMap[name] = v
		secretsNamesAws = append(secretsNamesAws, name)
	}

	chunks := utils.Chunkfy(secretsNamesAws, 20)
	self.logger.Debug().Any("chunks", chunks).Msg("secrets in chunks")

	for _, chunk := range chunks {
		result, err := secretManagerClient.BatchGetSecretValue(context.Background(), &secretsmanager.BatchGetSecretValueInput{
			SecretIdList: chunk,
		})
		if err != nil {
			self.logger.Error().Err(err).Msg("fail to get secrets")
			return err
		}

		for _, v := range result.SecretValues {
			if v.Name == nil {
				self.logger.Error().Any("v", v).Msg("fail to get secret, \"Name\" is nil")
				return fmt.Errorf("fail to get secret")
			}
			if v.SecretString == nil {
				self.logger.Error().Msgf("fail to get secret \"%s\", \"SecretString\" is nil", *v.Name)
				return fmt.Errorf("secret \"%s\" has no \"SecretString\"", *v.Name)
			}

			localName := namesMap[*v.Name]

			switch localName {
			case "DatabaseUrl":
				self.secrets.DatabaseUrl = *v.SecretString
			case "GoogleClientSecret":
				self.secrets.GoogleClientSecret = *v.SecretString
			}
		}
	}

	self.logger.Trace().Msg("all secrets loaded")

	return nil
}
