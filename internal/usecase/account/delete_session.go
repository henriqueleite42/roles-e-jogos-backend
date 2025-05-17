package account_usecase

import (
	"context"

	account_repository "github.com/henriqueleite42/roles-e-jogos-backend/internal/repository/account"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/utils"
)

func (self *AccountUsecaseImplementation) DeleteSession(ctx context.Context, i *DeleteSessionInput) error {
	ctx, err := utils.SetDbInCtx(self.Db, ctx)
	if err != nil {
		return err
	}

	err = self.AccountRepository.DeleteSession(ctx, &account_repository.DeleteSessionInput{
		SessionId: i.SessionId,
	})
	if err != nil {
		self.Logger.Error().Err(err).Msg("fail to delete session")
		return err
	}

	return nil
}
