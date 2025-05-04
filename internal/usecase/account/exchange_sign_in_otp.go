package account_usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/models"
	account_repository "github.com/henriqueleite42/roles-e-jogos-backend/internal/repository/account"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/utils"
)

func (self *AccountUsecaseImplementation) ExchangeSignInOtp(ctx context.Context, i *ExchangeSignInOtpInput) (*models.AccountData, error) {
	ctx, err := utils.SetDbInCtx(self.Db, ctx)
	if err != nil {
		return nil, err
	}

	otp, err := self.AccountRepository.GetOtp(ctx, &account_repository.GetOtpInput{
		AccountId: i.AccountId,
		Code:      i.Otp,
		Purpose:   models.OtpPurpose_SignIn,
	})
	if err != nil {
		if err.Error() != "not found" {
			return nil, fmt.Errorf("invalid")
		}
		return nil, err
	}

	expirationDate := otp.CreatedAt.Add(15 * time.Minute)
	if otp.CreatedAt.After(expirationDate) {
		return nil, fmt.Errorf("expired")
	}

	account, err := self.AccountRepository.GetAccountDataById(ctx, &account_repository.GetAccountDataByIdInput{
		AccountId: i.AccountId,
	})
	if err != nil {
		return nil, err
	}

	return account, nil
}
